// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package command

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/shulutkov/yellow-pages/command/acl"
	aclagent "github.com/shulutkov/yellow-pages/command/acl/agenttokens"
	aclam "github.com/shulutkov/yellow-pages/command/acl/authmethod"
	aclamcreate "github.com/shulutkov/yellow-pages/command/acl/authmethod/create"
	aclamdelete "github.com/shulutkov/yellow-pages/command/acl/authmethod/delete"
	aclamlist "github.com/shulutkov/yellow-pages/command/acl/authmethod/list"
	aclamread "github.com/shulutkov/yellow-pages/command/acl/authmethod/read"
	aclamupdate "github.com/shulutkov/yellow-pages/command/acl/authmethod/update"
	aclbr "github.com/shulutkov/yellow-pages/command/acl/bindingrule"
	aclbrcreate "github.com/shulutkov/yellow-pages/command/acl/bindingrule/create"
	aclbrdelete "github.com/shulutkov/yellow-pages/command/acl/bindingrule/delete"
	aclbrlist "github.com/shulutkov/yellow-pages/command/acl/bindingrule/list"
	aclbrread "github.com/shulutkov/yellow-pages/command/acl/bindingrule/read"
	aclbrupdate "github.com/shulutkov/yellow-pages/command/acl/bindingrule/update"
	aclbootstrap "github.com/shulutkov/yellow-pages/command/acl/bootstrap"
	aclpolicy "github.com/shulutkov/yellow-pages/command/acl/policy"
	aclpcreate "github.com/shulutkov/yellow-pages/command/acl/policy/create"
	aclpdelete "github.com/shulutkov/yellow-pages/command/acl/policy/delete"
	aclplist "github.com/shulutkov/yellow-pages/command/acl/policy/list"
	aclpread "github.com/shulutkov/yellow-pages/command/acl/policy/read"
	aclpupdate "github.com/shulutkov/yellow-pages/command/acl/policy/update"
	aclrole "github.com/shulutkov/yellow-pages/command/acl/role"
	aclrcreate "github.com/shulutkov/yellow-pages/command/acl/role/create"
	aclrdelete "github.com/shulutkov/yellow-pages/command/acl/role/delete"
	aclrlist "github.com/shulutkov/yellow-pages/command/acl/role/list"
	aclrread "github.com/shulutkov/yellow-pages/command/acl/role/read"
	aclrupdate "github.com/shulutkov/yellow-pages/command/acl/role/update"
	acltoken "github.com/shulutkov/yellow-pages/command/acl/token"
	acltclone "github.com/shulutkov/yellow-pages/command/acl/token/clone"
	acltcreate "github.com/shulutkov/yellow-pages/command/acl/token/create"
	acltdelete "github.com/shulutkov/yellow-pages/command/acl/token/delete"
	acltlist "github.com/shulutkov/yellow-pages/command/acl/token/list"
	acltread "github.com/shulutkov/yellow-pages/command/acl/token/read"
	acltupdate "github.com/shulutkov/yellow-pages/command/acl/token/update"
	"github.com/shulutkov/yellow-pages/command/agent"
	"github.com/shulutkov/yellow-pages/command/catalog"
	catlistdc "github.com/shulutkov/yellow-pages/command/catalog/list/dc"
	catlistnodes "github.com/shulutkov/yellow-pages/command/catalog/list/nodes"
	catlistsvc "github.com/shulutkov/yellow-pages/command/catalog/list/services"
	"github.com/shulutkov/yellow-pages/command/config"
	configdelete "github.com/shulutkov/yellow-pages/command/config/delete"
	configlist "github.com/shulutkov/yellow-pages/command/config/list"
	configread "github.com/shulutkov/yellow-pages/command/config/read"
	configwrite "github.com/shulutkov/yellow-pages/command/config/write"
	"github.com/shulutkov/yellow-pages/command/connect"
	"github.com/shulutkov/yellow-pages/command/connect/ca"
	caget "github.com/shulutkov/yellow-pages/command/connect/ca/get"
	caset "github.com/shulutkov/yellow-pages/command/connect/ca/set"
	"github.com/shulutkov/yellow-pages/command/connect/envoy"
	pipebootstrap "github.com/shulutkov/yellow-pages/command/connect/envoy/pipe-bootstrap"
	"github.com/shulutkov/yellow-pages/command/connect/expose"
	"github.com/shulutkov/yellow-pages/command/connect/proxy"
	"github.com/shulutkov/yellow-pages/command/connect/redirecttraffic"
	"github.com/shulutkov/yellow-pages/command/debug"
	"github.com/shulutkov/yellow-pages/command/event"
	"github.com/shulutkov/yellow-pages/command/exec"
	"github.com/shulutkov/yellow-pages/command/forceleave"
	"github.com/shulutkov/yellow-pages/command/info"
	"github.com/shulutkov/yellow-pages/command/intention"
	ixncheck "github.com/shulutkov/yellow-pages/command/intention/check"
	ixncreate "github.com/shulutkov/yellow-pages/command/intention/create"
	ixndelete "github.com/shulutkov/yellow-pages/command/intention/delete"
	ixnget "github.com/shulutkov/yellow-pages/command/intention/get"
	ixnlist "github.com/shulutkov/yellow-pages/command/intention/list"
	ixnmatch "github.com/shulutkov/yellow-pages/command/intention/match"
	"github.com/shulutkov/yellow-pages/command/join"
	"github.com/shulutkov/yellow-pages/command/keygen"
	"github.com/shulutkov/yellow-pages/command/keyring"
	"github.com/shulutkov/yellow-pages/command/kv"
	kvdel "github.com/shulutkov/yellow-pages/command/kv/del"
	kvexp "github.com/shulutkov/yellow-pages/command/kv/exp"
	kvget "github.com/shulutkov/yellow-pages/command/kv/get"
	kvimp "github.com/shulutkov/yellow-pages/command/kv/imp"
	kvput "github.com/shulutkov/yellow-pages/command/kv/put"
	"github.com/shulutkov/yellow-pages/command/leave"
	"github.com/shulutkov/yellow-pages/command/lock"
	"github.com/shulutkov/yellow-pages/command/login"
	"github.com/shulutkov/yellow-pages/command/logout"
	"github.com/shulutkov/yellow-pages/command/maint"
	"github.com/shulutkov/yellow-pages/command/members"
	"github.com/shulutkov/yellow-pages/command/monitor"
	"github.com/shulutkov/yellow-pages/command/operator"
	operauto "github.com/shulutkov/yellow-pages/command/operator/autopilot"
	operautoget "github.com/shulutkov/yellow-pages/command/operator/autopilot/get"
	operautoset "github.com/shulutkov/yellow-pages/command/operator/autopilot/set"
	operautostate "github.com/shulutkov/yellow-pages/command/operator/autopilot/state"
	operraft "github.com/shulutkov/yellow-pages/command/operator/raft"
	operraftlist "github.com/shulutkov/yellow-pages/command/operator/raft/listpeers"
	operraftremove "github.com/shulutkov/yellow-pages/command/operator/raft/removepeer"
	"github.com/shulutkov/yellow-pages/command/operator/raft/transferleader"
	"github.com/shulutkov/yellow-pages/command/operator/usage"
	"github.com/shulutkov/yellow-pages/command/operator/usage/instances"
	"github.com/shulutkov/yellow-pages/command/peering"
	peerdelete "github.com/shulutkov/yellow-pages/command/peering/delete"
	peerestablish "github.com/shulutkov/yellow-pages/command/peering/establish"
	peergenerate "github.com/shulutkov/yellow-pages/command/peering/generate"
	peerlist "github.com/shulutkov/yellow-pages/command/peering/list"
	peerread "github.com/shulutkov/yellow-pages/command/peering/read"
	"github.com/shulutkov/yellow-pages/command/reload"
	"github.com/shulutkov/yellow-pages/command/rtt"
	"github.com/shulutkov/yellow-pages/command/services"
	svcsderegister "github.com/shulutkov/yellow-pages/command/services/deregister"
	svcsexport "github.com/shulutkov/yellow-pages/command/services/export"
	svcsregister "github.com/shulutkov/yellow-pages/command/services/register"
	"github.com/shulutkov/yellow-pages/command/snapshot"
	snapinspect "github.com/shulutkov/yellow-pages/command/snapshot/inspect"
	snaprestore "github.com/shulutkov/yellow-pages/command/snapshot/restore"
	snapsave "github.com/shulutkov/yellow-pages/command/snapshot/save"
	"github.com/shulutkov/yellow-pages/command/tls"
	tlsca "github.com/shulutkov/yellow-pages/command/tls/ca"
	tlscacreate "github.com/shulutkov/yellow-pages/command/tls/ca/create"
	tlscert "github.com/shulutkov/yellow-pages/command/tls/cert"
	tlscertcreate "github.com/shulutkov/yellow-pages/command/tls/cert/create"
	"github.com/shulutkov/yellow-pages/command/troubleshoot"
	troubleshootports "github.com/shulutkov/yellow-pages/command/troubleshoot/ports"
	troubleshootproxy "github.com/shulutkov/yellow-pages/command/troubleshoot/proxy"
	troubleshootupstreams "github.com/shulutkov/yellow-pages/command/troubleshoot/upstreams"
	"github.com/shulutkov/yellow-pages/command/validate"
	"github.com/shulutkov/yellow-pages/command/version"
	"github.com/shulutkov/yellow-pages/command/watch"

	mcli "github.com/mitchellh/cli"

	"github.com/shulutkov/yellow-pages/command/cli"
)

// RegisteredCommands returns a realized mapping of available CLI commands in a format that
// the CLI class can consume.
func RegisteredCommands(ui cli.Ui) map[string]mcli.CommandFactory {
	registry := map[string]mcli.CommandFactory{}
	registerCommands(ui, registry,
		entry{"acl", func(cli.Ui) (cli.Command, error) { return acl.New(), nil }},
		entry{"acl bootstrap", func(ui cli.Ui) (cli.Command, error) { return aclbootstrap.New(ui), nil }},
		entry{"acl policy", func(cli.Ui) (cli.Command, error) { return aclpolicy.New(), nil }},
		entry{"acl policy create", func(ui cli.Ui) (cli.Command, error) { return aclpcreate.New(ui), nil }},
		entry{"acl policy list", func(ui cli.Ui) (cli.Command, error) { return aclplist.New(ui), nil }},
		entry{"acl policy read", func(ui cli.Ui) (cli.Command, error) { return aclpread.New(ui), nil }},
		entry{"acl policy update", func(ui cli.Ui) (cli.Command, error) { return aclpupdate.New(ui), nil }},
		entry{"acl policy delete", func(ui cli.Ui) (cli.Command, error) { return aclpdelete.New(ui), nil }},
		entry{"acl set-agent-token", func(ui cli.Ui) (cli.Command, error) { return aclagent.New(ui), nil }},
		entry{"acl token", func(cli.Ui) (cli.Command, error) { return acltoken.New(), nil }},
		entry{"acl token create", func(ui cli.Ui) (cli.Command, error) { return acltcreate.New(ui), nil }},
		entry{"acl token clone", func(ui cli.Ui) (cli.Command, error) { return acltclone.New(ui), nil }},
		entry{"acl token list", func(ui cli.Ui) (cli.Command, error) { return acltlist.New(ui), nil }},
		entry{"acl token read", func(ui cli.Ui) (cli.Command, error) { return acltread.New(ui), nil }},
		entry{"acl token update", func(ui cli.Ui) (cli.Command, error) { return acltupdate.New(ui), nil }},
		entry{"acl token delete", func(ui cli.Ui) (cli.Command, error) { return acltdelete.New(ui), nil }},
		entry{"acl role", func(cli.Ui) (cli.Command, error) { return aclrole.New(), nil }},
		entry{"acl role create", func(ui cli.Ui) (cli.Command, error) { return aclrcreate.New(ui), nil }},
		entry{"acl role list", func(ui cli.Ui) (cli.Command, error) { return aclrlist.New(ui), nil }},
		entry{"acl role read", func(ui cli.Ui) (cli.Command, error) { return aclrread.New(ui), nil }},
		entry{"acl role update", func(ui cli.Ui) (cli.Command, error) { return aclrupdate.New(ui), nil }},
		entry{"acl role delete", func(ui cli.Ui) (cli.Command, error) { return aclrdelete.New(ui), nil }},
		entry{"acl auth-method", func(cli.Ui) (cli.Command, error) { return aclam.New(), nil }},
		entry{"acl auth-method create", func(ui cli.Ui) (cli.Command, error) { return aclamcreate.New(ui), nil }},
		entry{"acl auth-method list", func(ui cli.Ui) (cli.Command, error) { return aclamlist.New(ui), nil }},
		entry{"acl auth-method read", func(ui cli.Ui) (cli.Command, error) { return aclamread.New(ui), nil }},
		entry{"acl auth-method update", func(ui cli.Ui) (cli.Command, error) { return aclamupdate.New(ui), nil }},
		entry{"acl auth-method delete", func(ui cli.Ui) (cli.Command, error) { return aclamdelete.New(ui), nil }},
		entry{"acl binding-rule", func(cli.Ui) (cli.Command, error) { return aclbr.New(), nil }},
		entry{"acl binding-rule create", func(ui cli.Ui) (cli.Command, error) { return aclbrcreate.New(ui), nil }},
		entry{"acl binding-rule list", func(ui cli.Ui) (cli.Command, error) { return aclbrlist.New(ui), nil }},
		entry{"acl binding-rule read", func(ui cli.Ui) (cli.Command, error) { return aclbrread.New(ui), nil }},
		entry{"acl binding-rule update", func(ui cli.Ui) (cli.Command, error) { return aclbrupdate.New(ui), nil }},
		entry{"acl binding-rule delete", func(ui cli.Ui) (cli.Command, error) { return aclbrdelete.New(ui), nil }},
		entry{"agent", func(ui cli.Ui) (cli.Command, error) { return agent.New(ui), nil }},
		entry{"catalog", func(cli.Ui) (cli.Command, error) { return catalog.New(), nil }},
		entry{"catalog datacenters", func(ui cli.Ui) (cli.Command, error) { return catlistdc.New(ui), nil }},
		entry{"catalog nodes", func(ui cli.Ui) (cli.Command, error) { return catlistnodes.New(ui), nil }},
		entry{"catalog services", func(ui cli.Ui) (cli.Command, error) { return catlistsvc.New(ui), nil }},
		entry{"config", func(ui cli.Ui) (cli.Command, error) { return config.New(), nil }},
		entry{"config delete", func(ui cli.Ui) (cli.Command, error) { return configdelete.New(ui), nil }},
		entry{"config list", func(ui cli.Ui) (cli.Command, error) { return configlist.New(ui), nil }},
		entry{"config read", func(ui cli.Ui) (cli.Command, error) { return configread.New(ui), nil }},
		entry{"config write", func(ui cli.Ui) (cli.Command, error) { return configwrite.New(ui), nil }},
		entry{"connect", func(ui cli.Ui) (cli.Command, error) { return connect.New(), nil }},
		entry{"connect ca", func(ui cli.Ui) (cli.Command, error) { return ca.New(), nil }},
		entry{"connect ca get-config", func(ui cli.Ui) (cli.Command, error) { return caget.New(ui), nil }},
		entry{"connect ca set-config", func(ui cli.Ui) (cli.Command, error) { return caset.New(ui), nil }},
		entry{"connect proxy", func(ui cli.Ui) (cli.Command, error) { return proxy.New(ui, MakeShutdownCh()), nil }},
		entry{"connect envoy", func(ui cli.Ui) (cli.Command, error) { return envoy.New(ui), nil }},
		entry{"connect envoy pipe-bootstrap", func(ui cli.Ui) (cli.Command, error) { return pipebootstrap.New(ui), nil }},
		entry{"connect expose", func(ui cli.Ui) (cli.Command, error) { return expose.New(ui), nil }},
		entry{"connect redirect-traffic", func(ui cli.Ui) (cli.Command, error) { return redirecttraffic.New(ui), nil }},
		entry{"debug", func(ui cli.Ui) (cli.Command, error) { return debug.New(ui), nil }},
		entry{"event", func(ui cli.Ui) (cli.Command, error) { return event.New(ui), nil }},
		entry{"exec", func(ui cli.Ui) (cli.Command, error) { return exec.New(ui, MakeShutdownCh()), nil }},
		entry{"force-leave", func(ui cli.Ui) (cli.Command, error) { return forceleave.New(ui), nil }},
		entry{"info", func(ui cli.Ui) (cli.Command, error) { return info.New(ui), nil }},
		entry{"intention", func(ui cli.Ui) (cli.Command, error) { return intention.New(), nil }},
		entry{"intention check", func(ui cli.Ui) (cli.Command, error) { return ixncheck.New(ui), nil }},
		entry{"intention create", func(ui cli.Ui) (cli.Command, error) { return ixncreate.New(ui), nil }},
		entry{"intention delete", func(ui cli.Ui) (cli.Command, error) { return ixndelete.New(ui), nil }},
		entry{"intention get", func(ui cli.Ui) (cli.Command, error) { return ixnget.New(ui), nil }},
		entry{"intention list", func(ui cli.Ui) (cli.Command, error) { return ixnlist.New(ui), nil }},
		entry{"intention match", func(ui cli.Ui) (cli.Command, error) { return ixnmatch.New(ui), nil }},
		entry{"join", func(ui cli.Ui) (cli.Command, error) { return join.New(ui), nil }},
		entry{"keygen", func(ui cli.Ui) (cli.Command, error) { return keygen.New(ui), nil }},
		entry{"keyring", func(ui cli.Ui) (cli.Command, error) { return keyring.New(ui), nil }},
		entry{"kv", func(cli.Ui) (cli.Command, error) { return kv.New(), nil }},
		entry{"kv delete", func(ui cli.Ui) (cli.Command, error) { return kvdel.New(ui), nil }},
		entry{"kv export", func(ui cli.Ui) (cli.Command, error) { return kvexp.New(ui), nil }},
		entry{"kv get", func(ui cli.Ui) (cli.Command, error) { return kvget.New(ui), nil }},
		entry{"kv import", func(ui cli.Ui) (cli.Command, error) { return kvimp.New(ui), nil }},
		entry{"kv put", func(ui cli.Ui) (cli.Command, error) { return kvput.New(ui), nil }},
		entry{"leave", func(ui cli.Ui) (cli.Command, error) { return leave.New(ui), nil }},
		entry{"lock", func(ui cli.Ui) (cli.Command, error) { return lock.New(ui, MakeShutdownCh()), nil }},
		entry{"login", func(ui cli.Ui) (cli.Command, error) { return login.New(ui), nil }},
		entry{"logout", func(ui cli.Ui) (cli.Command, error) { return logout.New(ui), nil }},
		entry{"maint", func(ui cli.Ui) (cli.Command, error) { return maint.New(ui), nil }},
		entry{"members", func(ui cli.Ui) (cli.Command, error) { return members.New(ui), nil }},
		entry{"monitor", func(ui cli.Ui) (cli.Command, error) { return monitor.New(ui, MakeShutdownCh()), nil }},
		entry{"operator", func(cli.Ui) (cli.Command, error) { return operator.New(), nil }},
		entry{"operator autopilot", func(cli.Ui) (cli.Command, error) { return operauto.New(), nil }},
		entry{"operator autopilot get-config", func(ui cli.Ui) (cli.Command, error) { return operautoget.New(ui), nil }},
		entry{"operator autopilot set-config", func(ui cli.Ui) (cli.Command, error) { return operautoset.New(ui), nil }},
		entry{"operator autopilot state", func(ui cli.Ui) (cli.Command, error) { return operautostate.New(ui), nil }},
		entry{"operator raft", func(cli.Ui) (cli.Command, error) { return operraft.New(), nil }},
		entry{"operator raft list-peers", func(ui cli.Ui) (cli.Command, error) { return operraftlist.New(ui), nil }},
		entry{"operator raft remove-peer", func(ui cli.Ui) (cli.Command, error) { return operraftremove.New(ui), nil }},
		entry{"operator raft transfer-leader", func(ui cli.Ui) (cli.Command, error) { return transferleader.New(ui), nil }},
		entry{"operator usage", func(ui cli.Ui) (cli.Command, error) { return usage.New(), nil }},
		entry{"operator usage instances", func(ui cli.Ui) (cli.Command, error) { return instances.New(ui), nil }},
		entry{"peering", func(cli.Ui) (cli.Command, error) { return peering.New(), nil }},
		entry{"peering delete", func(ui cli.Ui) (cli.Command, error) { return peerdelete.New(ui), nil }},
		entry{"peering generate-token", func(ui cli.Ui) (cli.Command, error) { return peergenerate.New(ui), nil }},
		entry{"peering establish", func(ui cli.Ui) (cli.Command, error) { return peerestablish.New(ui), nil }},
		entry{"peering list", func(ui cli.Ui) (cli.Command, error) { return peerlist.New(ui), nil }},
		entry{"peering read", func(ui cli.Ui) (cli.Command, error) { return peerread.New(ui), nil }},
		entry{"reload", func(ui cli.Ui) (cli.Command, error) { return reload.New(ui), nil }},
		entry{"rtt", func(ui cli.Ui) (cli.Command, error) { return rtt.New(ui), nil }},
		entry{"services", func(cli.Ui) (cli.Command, error) { return services.New(), nil }},
		entry{"services register", func(ui cli.Ui) (cli.Command, error) { return svcsregister.New(ui), nil }},
		entry{"services deregister", func(ui cli.Ui) (cli.Command, error) { return svcsderegister.New(ui), nil }},
		entry{"services export", func(ui cli.Ui) (cli.Command, error) { return svcsexport.New(ui), nil }},
		entry{"snapshot", func(cli.Ui) (cli.Command, error) { return snapshot.New(), nil }},
		entry{"snapshot inspect", func(ui cli.Ui) (cli.Command, error) { return snapinspect.New(ui), nil }},
		entry{"snapshot restore", func(ui cli.Ui) (cli.Command, error) { return snaprestore.New(ui), nil }},
		entry{"snapshot save", func(ui cli.Ui) (cli.Command, error) { return snapsave.New(ui), nil }},
		entry{"tls", func(ui cli.Ui) (cli.Command, error) { return tls.New(), nil }},
		entry{"tls ca", func(ui cli.Ui) (cli.Command, error) { return tlsca.New(), nil }},
		entry{"tls ca create", func(ui cli.Ui) (cli.Command, error) { return tlscacreate.New(ui), nil }},
		entry{"tls cert", func(ui cli.Ui) (cli.Command, error) { return tlscert.New(), nil }},
		entry{"tls cert create", func(ui cli.Ui) (cli.Command, error) { return tlscertcreate.New(ui), nil }},
		entry{"troubleshoot", func(ui cli.Ui) (cli.Command, error) { return troubleshoot.New(), nil }},
		entry{"troubleshoot proxy", func(ui cli.Ui) (cli.Command, error) { return troubleshootproxy.New(ui), nil }},
		entry{"troubleshoot upstreams", func(ui cli.Ui) (cli.Command, error) { return troubleshootupstreams.New(ui), nil }},
		entry{"troubleshoot ports", func(ui cli.Ui) (cli.Command, error) { return troubleshootports.New(ui), nil }},
		entry{"validate", func(ui cli.Ui) (cli.Command, error) { return validate.New(ui), nil }},
		entry{"version", func(ui cli.Ui) (cli.Command, error) { return version.New(ui), nil }},
		entry{"watch", func(ui cli.Ui) (cli.Command, error) { return watch.New(ui, MakeShutdownCh()), nil }},
	)
	registerEnterpriseCommands(ui, registry)
	return registry
}

// factory is a function that returns a new instance of a CLI-sub command.
type factory func(cli.Ui) (cli.Command, error)

// entry is a struct that contains a command's name and a factory for that command.
type entry struct {
	name string
	fn   factory
}

func registerCommands(ui cli.Ui, m map[string]mcli.CommandFactory, cmdEntries ...entry) {
	for _, ent := range cmdEntries {
		thisFn := ent.fn
		if _, ok := m[ent.name]; ok {
			panic(fmt.Sprintf("duplicate command: %q", ent.name))
		}
		m[ent.name] = func() (mcli.Command, error) {
			return thisFn(ui)
		}
	}
}

// MakeShutdownCh returns a channel that can be used for shutdown notifications
// for commands. This channel will send a message for every interrupt or SIGTERM
// received.
// Deprecated: use signal.NotifyContext
func MakeShutdownCh() <-chan struct{} {
	resultCh := make(chan struct{})
	signalCh := make(chan os.Signal, 4)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		for {
			<-signalCh
			resultCh <- struct{}{}
		}
	}()

	return resultCh
}
