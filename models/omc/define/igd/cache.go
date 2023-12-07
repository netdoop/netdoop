package igd

import (
	"sync"
)

var defaultCacheManager *CacheManager
var defaultCacheManagerOnce sync.Once

func GetDefaultCacheManager() *CacheManager {
	defaultCacheManagerOnce.Do(func() {
		defaultCacheManager = NewCacheManager()
	})
	return defaultCacheManager
}

type CacheManager struct {
	cache sync.Map
}

func NewCacheManager() *CacheManager {
	return &CacheManager{
		cache: sync.Map{},
	}
}

func (m *CacheManager) Get(id string) *InternetGatewayDevice {
	if obj, ok := m.cache.Load(id); ok {
		return obj.(*InternetGatewayDevice)
	}
	return nil
}

func (m *CacheManager) Set(id string, obj *InternetGatewayDevice) {
	m.cache.Store(id, obj)
}

func (m *CacheManager) Update(id string, values map[string]string) error {
	obj := m.Get(id)
	if obj == nil {
		obj = &InternetGatewayDevice{}
	}
	obj.ReadValues(values)
	m.cache.Store(id, obj)
	return nil
}
