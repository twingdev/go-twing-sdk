package common

import (
	"context"
	"github.com/eko/gocache/lib/v4/cache"
	_ "github.com/eko/gocache/lib/v4/cache"
	gocache_store "github.com/eko/gocache/store/go_cache/v4"
	gocache "github.com/patrickmn/go-cache"
	"time"
)

type InMemoryCache struct {
	manager   *cache.Cache[[]byte]
	memory    *gocache_store.GoCacheStore
	memClient gocache_store.GoCacheClientInterface
}

func NewInMemoryCache() *InMemoryCache {

	imc := &InMemoryCache{}
	imc.memClient = gocache.New(time.Second*3600, time.Second*3600*24)
	imc.memory = gocache_store.NewGoCache(imc.memClient)
	imc.manager = cache.New[[]byte](imc.memory)

	return imc

}

func (imc *InMemoryCache) Get(ctx context.Context, k string) ([]byte, error) {
	value, err := imc.manager.Get(ctx, k)
	return value, err
}

func (imc *InMemoryCache) Set(ctx context.Context, k string, v []byte) error {
	return imc.manager.Set(ctx, k, v)
}

func (imc *InMemoryCache) Delete(ctx context.Context, k string) error {
	return imc.manager.Delete(ctx, k)
}

func (imc *InMemoryCache) ClearAll(ctx context.Context) error {
	return imc.manager.Clear(ctx)
}

var TwingCache *InMemoryCache
var LocalCache *InMemoryCache

func init() {
	TwingCache = NewInMemoryCache()
	LocalCache = NewInMemoryCache()
}
