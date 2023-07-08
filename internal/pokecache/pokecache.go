package pokecache

import (
	"errors"
	"sync"
	"time"
)

type CacheEntry struct {
	Value     string
	createdAt time.Time
}

type CacheMap struct {
	cache map[string]CacheEntry
	mutex sync.Mutex
}

func (cacheMap *CacheMap) NewCache(cacheDuration time.Duration) {
	cache := make(map[string]CacheEntry)
	cacheMap.cache = cache
	go readLoop(cacheDuration, cacheMap)
}

func readLoop(cacheDuration time.Duration, cacheMap *CacheMap) {

	for {
		for key, value := range cacheMap.cache {
			createdTime := value.createdAt
			currentTime := time.Now()

			expirationTime := createdTime.Add(cacheDuration)

			if expirationTime.Before(currentTime) {
				cacheMap.mutex.Lock()
				delete(cacheMap.cache, key)
				cacheMap.mutex.Unlock()
			}
		}
	}
}

func (cacheMap *CacheMap) Add(key string, value string) {
	currentTime := time.Now()
	cacheMap.mutex.Lock()
	cacheMap.cache[key] = CacheEntry{value, currentTime}
	cacheMap.mutex.Unlock()
}

func (cacheMap *CacheMap) Get(key string) (CacheEntry, error) {
	value, ok := cacheMap.cache[key]
	if !ok {
		return CacheEntry{}, errors.New("Key not found or might have expired")
	}
	return value, nil
}
