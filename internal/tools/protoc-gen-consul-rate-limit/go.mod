module github.com/arenadata/consul/internal/tools/protoc-gen-consul-rate-limit

go 1.24.0

replace github.com/arenadata/consul/proto-public => ../../../proto-public

require (
	github.com/arenadata/consul/proto-public v0.0.0
	google.golang.org/protobuf v1.34.2
)

require github.com/google/go-cmp v0.5.9 // indirect
