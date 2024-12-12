package cache

import (
	"sync"
)

type cacheOperationOptions struct {
}

var operationOptionsPool = &sync.Pool{
	New: func() interface{} {
		return &cacheOperationOptions{}
	},
}

func (p *cacheOperationOptions) reset() {
}

func newCacheOperationOptions() *cacheOperationOptions {
	option, _ := operationOptionsPool.Get().(*cacheOperationOptions)
	option.reset()
	return option
}

func recycleCacheOperationOptions(option interface{}) {
	operationOptionsPool.Put(option)
}

type OperationOption func(*cacheOperationOptions)
