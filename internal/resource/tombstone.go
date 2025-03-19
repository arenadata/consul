package resource

import "github.com/arenadata/consul/proto-public/pbresource"

var (
	TypeV1Tombstone = &pbresource.Type{
		Group:        "internal",
		GroupVersion: "v1",
		Kind:         "Tombstone",
	}
)
