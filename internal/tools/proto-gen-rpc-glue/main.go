// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	flagPath = flag.String("path", "", "path of file to load")
	verbose  = flag.Bool("v", false, "verbose output")
)

const (
	annotationPrefix = "@consul-rpc-glue:"
	outputFileSuffix = ".rpcglue.pb.go"
)

func main() {
	flag.Parse()

	log.SetFlags(0)

	if *flagPath == "" {
		log.Fatal("missing required -path argument")
	}

	if err := run(*flagPath); err != nil {
		log.Fatal(err)
	}
}

func run(path string) error {
	fi, err := os.Stat(path)
	if err != nil {
		return err
	}

	if fi.IsDir() {
		return fmt.Errorf("argument must be a file: %s", path)
	}

	if !strings.HasSuffix(path, ".pb.go") {
		return fmt.Errorf("file must end with .pb.go: %s", path)
	}

	if err := processFile(path); err != nil {
		return fmt.Errorf("error processing file %q: %v", path, err)
	}

	return nil
}

func processFile(path string) error {
	if *verbose {
		log.Printf("visiting file %q", path)
	}

	fset := token.NewFileSet()
	tree, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	v := visitor{}
	ast.Walk(&v, tree)
	if err := v.Err(); err != nil {
		return err
	}

	if len(v.Types) == 0 {
		return nil
	}

	if *verbose {
		log.Printf("Package: %s", v.Package)
		log.Printf("BuildTags: %v", v.BuildTags)
		log.Println()
		for _, typ := range v.Types {
			log.Printf("Type: %s", typ.Name)
			ann := typ.Annotation
			if ann.ReadRequest != "" {
				log.Printf("    ReadRequest from %s", ann.ReadRequest)
			}
			if ann.WriteRequest != "" {
				log.Printf("    WriteRequest from %s", ann.WriteRequest)
			}
			if ann.TargetDatacenter != "" {
				log.Printf("    TargetDatacenter from %s", ann.TargetDatacenter)
			}
			if ann.QueryOptions != "" {
				log.Printf("    QueryOptions from %s", ann.QueryOptions)
			}
			if ann.QueryMeta != "" {
				log.Printf("    QueryMeta from %s", ann.QueryMeta)
			}
			if ann.Datacenter != "" {
				log.Printf("    Datacenter from %s", ann.Datacenter)
			}
		}
	}

	// generate output

	var buf bytes.Buffer

	if len(v.BuildTags) > 0 {
		for _, line := range v.BuildTags {
			buf.WriteString(line + "\n")
		}
		buf.WriteString("\n")
	}
	buf.WriteString("// Code generated by proto-gen-rpc-glue. DO NOT EDIT.\n\n")
	buf.WriteString("package " + v.Package + "\n")
	buf.WriteString(`
import (
	"time"

	"github.com/arenadata/consul/agent/structs"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ structs.RPCInfo
var _ time.Month

`)
	for _, typ := range v.Types {
		if typ.Annotation.WriteRequest != "" {
			buf.WriteString(fmt.Sprintf(tmplWriteRequest, typ.Name, typ.Annotation.WriteRequest))
		}
		if typ.Annotation.ReadRequest != "" {
			buf.WriteString(fmt.Sprintf(tmplReadRequest, typ.Name, typ.Annotation.ReadRequest))
		}
		if typ.Annotation.TargetDatacenter != "" {
			buf.WriteString(fmt.Sprintf(tmplTargetDatacenter, typ.Name, typ.Annotation.TargetDatacenter))
		}
		if typ.Annotation.QueryOptions != "" {
			buf.WriteString(fmt.Sprintf(tmplQueryOptions, typ.Name, typ.Annotation.QueryOptions))
		}
		if typ.Annotation.QueryMeta != "" {
			buf.WriteString(fmt.Sprintf(tmplQueryMeta, typ.Name, typ.Annotation.QueryMeta))
		}
		if typ.Annotation.Datacenter != "" {
			buf.WriteString(fmt.Sprintf(tmplDatacenter, typ.Name, typ.Annotation.Datacenter))
		}
	}

	// write to disk
	outFile := strings.TrimSuffix(path, ".pb.go") + outputFileSuffix
	if err := os.WriteFile(outFile, buf.Bytes(), 0644); err != nil {
		return err
	}

	// clean up
	cmd := exec.Command("gofmt", "-s", "-w", outFile)
	cmd.Stdout = nil
	cmd.Stderr = os.Stderr
	cmd.Stdin = nil
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error running 'gofmt -s -w %q': %v", outFile, err)
	}

	return nil
}

type TypeInfo struct {
	Name       string
	Annotation Annotation
}

type visitor struct {
	Package   string
	BuildTags []string
	Types     []TypeInfo
	Errs      []error
}

func (v *visitor) Err() error {
	switch len(v.Errs) {
	case 0:
		return nil
	case 1:
		return v.Errs[0]
	default:
		//
		var s []string
		for _, e := range v.Errs {
			s = append(s, e.Error())
		}
		return errors.New(strings.Join(s, "; "))
	}
}

var _ ast.Visitor = (*visitor)(nil)

func (v *visitor) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return v
	}

	switch x := node.(type) {
	case *ast.File:
		v.Package = x.Name.Name
		v.BuildTags = getRawBuildTags(x)
		for _, d := range x.Decls {
			gd, ok := d.(*ast.GenDecl)
			if !ok {
				continue
			}

			if gd.Doc == nil {
				continue
			} else if len(gd.Specs) != 1 {
				continue
			}
			spec := gd.Specs[0]

			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			ann, err := getAnnotation(gd.Doc.List)
			if err != nil {
				v.Errs = append(v.Errs, err)
				continue
			} else if ann.IsZero() {
				continue
			}

			v.Types = append(v.Types, TypeInfo{
				Name:       typeSpec.Name.Name,
				Annotation: ann,
			})

		}
	}
	return v
}

type Annotation struct {
	QueryMeta        string
	QueryOptions     string
	ReadRequest      string
	WriteRequest     string
	TargetDatacenter string
	Datacenter       string
	ReadTODO         string
	LeaderReadTODO   string
	WriteTODO        string
}

func (a Annotation) IsZero() bool {
	return a == Annotation{}
}

func getAnnotation(doc []*ast.Comment) (Annotation, error) {
	raw, ok := getRawStructAnnotation(doc)
	if !ok {
		return Annotation{}, nil
	}

	var ann Annotation

	parts := strings.Split(raw, ",")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		switch {
		case part == "ReadRequest":
			ann.ReadRequest = "ReadRequest"
		case strings.HasPrefix(part, "ReadRequest="):
			ann.ReadRequest = strings.TrimPrefix(part, "ReadRequest=")

		case part == "WriteRequest":
			ann.WriteRequest = "WriteRequest"
		case strings.HasPrefix(part, "WriteRequest="):
			ann.WriteRequest = strings.TrimPrefix(part, "WriteRequest=")

		case part == "TargetDatacenter":
			ann.TargetDatacenter = "TargetDatacenter"
		case strings.HasPrefix(part, "TargetDatacenter="):
			ann.TargetDatacenter = strings.TrimPrefix(part, "TargetDatacenter=")

		case part == "QueryOptions":
			ann.QueryOptions = "QueryOptions"
		case strings.HasPrefix(part, "QueryOptions="):
			ann.QueryOptions = strings.TrimPrefix(part, "QueryOptions=")

		case part == "QueryMeta":
			ann.QueryMeta = "QueryMeta"
		case strings.HasPrefix(part, "QueryMeta="):
			ann.QueryMeta = strings.TrimPrefix(part, "QueryMeta=")

		case part == "Datacenter":
			ann.Datacenter = "Datacenter"
		case strings.HasPrefix(part, "Datacenter="):
			ann.Datacenter = strings.TrimPrefix(part, "Datacenter=")

		default:
			return Annotation{}, fmt.Errorf("unexpected annotation part: %s", part)
		}
	}

	return ann, nil
}

func getRawStructAnnotation(doc []*ast.Comment) (string, bool) {
	for _, line := range doc {
		text := strings.TrimSpace(strings.TrimLeft(line.Text, "/"))

		ann := strings.TrimSpace(strings.TrimPrefix(text, annotationPrefix))

		if text != ann {
			return ann, true
		}
	}
	return "", false
}

func getRawBuildTags(file *ast.File) []string {
	// build tags are always the first group, at the very top
	if len(file.Comments) == 0 {
		return nil
	}
	cg := file.Comments[0]

	var out []string
	for _, line := range cg.List {
		text := strings.TrimSpace(strings.TrimLeft(line.Text, "/"))

		if !strings.HasPrefix(text, "go:build ") && !strings.HasPrefix(text, "+build") {
			break // stop at first non-build-tag
		}

		out = append(out, line.Text)
	}

	return out
}

const tmplWriteRequest = `
// AllowStaleRead implements structs.RPCInfo
func (msg *%[1]s) AllowStaleRead() bool {
	return false
}

// HasTimedOut implements structs.RPCInfo
func (msg *%[1]s) HasTimedOut(start time.Time, rpcHoldTimeout time.Duration, a time.Duration, b time.Duration) (bool, error) {
	if msg == nil || msg.%[2]s == nil {
		return false, nil
	}
	return msg.%[2]s.HasTimedOut(start, rpcHoldTimeout, a, b)
}

// IsRead implements structs.RPCInfo
func (msg *%[1]s) IsRead() bool {
	return false
}

// SetTokenSecret implements structs.RPCInfo
func (msg *%[1]s) SetTokenSecret(s string) {
    // TODO: initialize if nil
	msg.%[2]s.SetTokenSecret(s)
}

// TokenSecret implements structs.RPCInfo
func (msg *%[1]s) TokenSecret() string {
	if msg == nil || msg.%[2]s == nil {
		return ""
	}
	return msg.%[2]s.TokenSecret()
}

// Token implements structs.RPCInfo
func (msg *%[1]s) Token() string {
	if msg.%[2]s == nil {
		return ""
	}
	return msg.%[2]s.Token
}
`

const tmplReadRequest = `
// IsRead implements structs.RPCInfo
func (msg *%[1]s) IsRead() bool {
	return true
}

// AllowStaleRead implements structs.RPCInfo
func (msg *%[1]s) AllowStaleRead() bool {
    // TODO: initialize if nil
	return msg.%[2]s.AllowStaleRead()
}

// HasTimedOut implements structs.RPCInfo
func (msg *%[1]s) HasTimedOut(start time.Time, rpcHoldTimeout time.Duration, a time.Duration, b time.Duration) (bool, error) {
	if msg == nil || msg.%[2]s == nil {
		return false, nil
	}
	return msg.%[2]s.HasTimedOut(start, rpcHoldTimeout, a, b)
}

// SetTokenSecret implements structs.RPCInfo
func (msg *%[1]s) SetTokenSecret(s string) {
    // TODO: initialize if nil
	msg.%[2]s.SetTokenSecret(s)
}

// TokenSecret implements structs.RPCInfo
func (msg *%[1]s) TokenSecret() string {
	if msg == nil || msg.%[2]s == nil {
		return ""
	}
	return msg.%[2]s.TokenSecret()
}

// Token implements structs.RPCInfo
func (msg *%[1]s) Token() string {
	if msg.%[2]s == nil {
		return ""
	}
	return msg.%[2]s.Token
}
`

const tmplTargetDatacenter = `
// RequestDatacenter implements structs.RPCInfo
func (msg *%[1]s) RequestDatacenter() string {
	if msg == nil || msg.%[2]s == nil {
		return ""
	}
	return msg.%[2]s.GetDatacenter()
}
`

const tmplDatacenter = `
// RequestDatacenter implements structs.RPCInfo
func (msg *%[1]s) RequestDatacenter() string {
	if msg == nil {
		return ""
	}
	return msg.Datacenter
}
`

const tmplQueryOptions = `
// IsRead implements structs.RPCInfo
func (msg *%[1]s) IsRead() bool {
	return true
}

// AllowStaleRead implements structs.RPCInfo
func (msg *%[1]s) AllowStaleRead() bool {
	return msg.%[2]s.AllowStaleRead()
}

// BlockingTimeout implements pool.BlockableQuery
func (msg *%[1]s) BlockingTimeout(maxQueryTime, defaultQueryTime time.Duration) time.Duration {
	maxTime := structs.DurationFromProto(msg.%[2]s.GetMaxQueryTime())
	o := structs.QueryOptions{
		MaxQueryTime:  maxTime,
		MinQueryIndex: msg.%[2]s.GetMinQueryIndex(),
	}
	return o.BlockingTimeout(maxQueryTime, defaultQueryTime)
}

// HasTimedOut implements structs.RPCInfo
func (msg *%[1]s) HasTimedOut(start time.Time, rpcHoldTimeout time.Duration, a time.Duration, b time.Duration) (bool, error) {
	if msg == nil || msg.%[2]s == nil {
		return false, nil
	}
	return msg.%[2]s.HasTimedOut(start, rpcHoldTimeout, a, b)
}

// SetTokenSecret implements structs.RPCInfo
func (msg *%[1]s) SetTokenSecret(s string) {
    // TODO: initialize if nil
	msg.%[2]s.SetTokenSecret(s)
}

// TokenSecret implements structs.RPCInfo
func (msg *%[1]s) TokenSecret() string {
	if msg == nil || msg.%[2]s == nil {
		return ""
	}
	return msg.%[2]s.TokenSecret()
}

// Token implements structs.RPCInfo
func (msg *%[1]s) Token() string {
	if msg.%[2]s == nil {
		return ""
	}
	return msg.%[2]s.Token
}
// GetToken is required to implement blockingQueryOptions
func (msg *%[1]s) GetToken() string {
	if msg == nil || msg.%[2]s == nil {
		return ""
	}

	return msg.%[2]s.GetToken()
}
// GetMinQueryIndex is required to implement blockingQueryOptions
func (msg *%[1]s) GetMinQueryIndex() uint64 {
	if msg == nil || msg.%[2]s == nil {
		return 0
	}

	return msg.%[2]s.GetMinQueryIndex()
}
// GetMaxQueryTime is required to implement blockingQueryOptions
func (msg *%[1]s) GetMaxQueryTime() (time.Duration, error) {
	if msg == nil || msg.%[2]s == nil {
		return 0, nil
	}

	return structs.DurationFromProto(msg.%[2]s.GetMaxQueryTime()), nil
}

// GetRequireConsistent is required to implement blockingQueryOptions
func (msg *%[1]s) GetRequireConsistent() bool {
	if msg == nil || msg.%[2]s == nil {
		return false
	}
	return msg.%[2]s.RequireConsistent
}
`

const tmplQueryMeta = `
// SetLastContact is required to implement blockingQueryResponseMeta
func (msg *%[1]s) SetLastContact(d time.Duration) {
	if msg == nil || msg.%[2]s == nil {
		return
	}
	msg.%[2]s.SetLastContact(d)
}

// SetKnownLeader is required to implement blockingQueryResponseMeta
func (msg *%[1]s) SetKnownLeader(b bool) {
	if msg == nil || msg.%[2]s == nil {
		return
	}
	msg.%[2]s.SetKnownLeader(b)
}

// GetIndex is required to implement blockingQueryResponseMeta
func (msg *%[1]s) GetIndex() uint64 {
	if msg == nil || msg.%[2]s == nil {
		return 0
	}
	return msg.%[2]s.GetIndex()
}

// SetIndex is required to implement blockingQueryResponseMeta
func (msg *%[1]s) SetIndex(i uint64) {
	if msg == nil || msg.%[2]s == nil {
		return
	}
	msg.%[2]s.SetIndex(i)
}

// SetResultsFilteredByACLs is required to implement blockingQueryResponseMeta
func (msg *%[1]s) SetResultsFilteredByACLs(b bool) {
	if msg == nil || msg.%[2]s == nil {
		return
	}
	msg.%[2]s.SetResultsFilteredByACLs(b)
}
`
