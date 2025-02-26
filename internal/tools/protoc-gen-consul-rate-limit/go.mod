module github.com/shulutkov/yellow-pages/internal/tools/protoc-gen-consul-rate-limit

go 1.24.0

replace github.com/shulutkov/yellow-pages/proto-public => ../../../proto-public

require (
	github.com/shulutkov/yellow-pages/proto-public v0.0.0
	google.golang.org/protobuf v1.34.2
)

require github.com/google/go-cmp v0.5.9 // indirect
