package jwks

import (
	"sync"
	"time"
)

type Cache struct {
	mu   sync.RWMutex
	data JWKS
	exp  time.Time
}

var JWKSCache = &Cache{}

func (c *Cache) Get() (JWKS, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if time.Now().After(c.exp) {
		return JWKS{}, false
	}
	return c.data, true
}

func (c *Cache) Set(jwks JWKS) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data = jwks
	c.exp = time.Now().Add(5 * time.Minute)
}
