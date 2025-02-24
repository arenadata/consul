package resource

import "github.com/shulutkov/yellow-pages/proto-public/pbresource"

var (
	TypeV1Tombstone = &pbresource.Type{
		Group:        "internal",
		GroupVersion: "v1",
		Kind:         "Tombstone",
	}
)
