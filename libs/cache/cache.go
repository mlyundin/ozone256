package memorycache

import (
	"errors"
	"sync"

	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	namespace      = "route256"
	cacheSubsystem = "cache"

	HitCounter = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: cacheSubsystem,
		Name:      "hit_total",
	})

	MissCounter = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: cacheSubsystem,
		Name:      "miss_total",
	})

	HistogramResponseTime = promauto.NewHistogram(prometheus.HistogramOpts{
		Namespace: namespace,
		Subsystem: cacheSubsystem,
		Name:      "histogram_response_time_seconds",
		Buckets:   prometheus.ExponentialBuckets(0.0001, 2, 16),
	},
	)
)

// Cache struct cache
type Cache[Key comparable, Value any] struct {
	lock              sync.RWMutex
	items             map[Key]item[Value]
	defaultExpiration time.Duration
	cleanupInterval   time.Duration
}

// item struct cache item
type item[Value any] struct {
	value      Value
	expiration int64
}

// New. Initializing a new memory cache
func New[Key comparable, Value any](defaultExpiration, cleanupInterval time.Duration) *Cache[Key, Value] {

	// cache item
	cache := Cache[Key, Value]{
		items:             make(map[Key]item[Value]),
		defaultExpiration: defaultExpiration,
		cleanupInterval:   cleanupInterval,
	}

	if cleanupInterval > 0 {
		cache.startGC()
	}

	return &cache
}

// Set setting a cache by key
func (c *Cache[Key, Value]) Set(key Key, value Value, duration time.Duration) {

	var expiration int64

	if duration == 0 {
		duration = c.defaultExpiration
	}

	if duration > 0 {
		expiration = time.Now().Add(duration).UnixNano()
	}

	c.lock.Lock()

	defer c.lock.Unlock()

	c.items[key] = item[Value]{
		value:      value,
		expiration: expiration,
	}
}

// Get getting a cache by key
func (c *Cache[Key, Value]) Get(key Key) (Value, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	timeStart := time.Now()
	item, found := c.items[key]
	elapsed := time.Since(timeStart)
	HistogramResponseTime.Observe(elapsed.Seconds())

	if found && time.Now().UnixNano() <= item.expiration {
		HitCounter.Inc()
		return item.value, true
	}

	MissCounter.Inc()

	return item.value, false
}

// Delete cache by key
// Return false if key not found
func (c *Cache[Key, Value]) Delete(key Key) error {

	c.lock.Lock()

	defer c.lock.Unlock()

	if _, found := c.items[key]; !found {
		return errors.New("key not found")
	}

	delete(c.items, key)

	return nil
}

// startGC start Garbage Collection
func (c *Cache[Key, Value]) startGC() {
	go func() {
		for {
			<-time.After(c.cleanupInterval)

			if keys := c.expiredKeys(); len(keys) != 0 {
				c.clearItems(keys)
			}
		}
	}()
}

// expiredKeys returns key list which are expired.
func (c *Cache[Key, Value]) expiredKeys() (keys []Key) {

	c.lock.RLock()

	defer c.lock.RUnlock()

	for k, i := range c.items {
		if time.Now().UnixNano() > i.expiration && i.expiration > 0 {
			keys = append(keys, k)
		}
	}

	return
}

// clearItems removes all the items which key in keys.
func (c *Cache[Key, Value]) clearItems(keys []Key) {

	c.lock.Lock()

	defer c.lock.Unlock()

	for _, k := range keys {
		delete(c.items, k)
	}
}
