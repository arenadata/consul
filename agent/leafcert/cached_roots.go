package leafcert

import (
	"context"
	"errors"

	"github.com/arenadata/consul/agent/cache"
	cachetype "github.com/arenadata/consul/agent/cache-types"
	"github.com/arenadata/consul/agent/structs"
)

// NewCachedRootsReader returns a RootsReader that sources data from the agent cache.
func NewCachedRootsReader(cache *cache.Cache, dc string) RootsReader {
	return &agentCacheRootsReader{
		cache:      cache,
		datacenter: dc,
	}
}

type agentCacheRootsReader struct {
	cache      *cache.Cache
	datacenter string
}

var _ RootsReader = (*agentCacheRootsReader)(nil)

func (r *agentCacheRootsReader) Get() (*structs.IndexedCARoots, error) {
	// Background is fine here because this isn't a blocking query as no index is set.
	// Therefore this will just either be a cache hit or return once the non-blocking query returns.
	rawRoots, _, err := r.cache.Get(context.Background(), cachetype.ConnectCARootName, &structs.DCSpecificRequest{
		Datacenter: r.datacenter,
	})
	if err != nil {
		return nil, err
	}
	roots, ok := rawRoots.(*structs.IndexedCARoots)
	if !ok {
		return nil, errors.New("invalid RootCA response type")
	}
	return roots, nil
}

func (r *agentCacheRootsReader) Notify(ctx context.Context, correlationID string, ch chan<- cache.UpdateEvent) error {
	return r.cache.Notify(ctx, cachetype.ConnectCARootName, &structs.DCSpecificRequest{
		Datacenter: r.datacenter,
	}, correlationID, ch)
}
