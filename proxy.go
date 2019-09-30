package readthrough

// Proxy behaves as a read-through cache for an origin (DB, API server, etc)
type Proxy interface {
	Get(key string) (bool, interface{}, error)
	Set(key string, val interface{}) error
}

// NewInMemoryProxy creates an InMemory Proxy object
func NewInMemoryProxy() Proxy {
	// TODO: implement an eviction algorithm
	return &InMemoryProxy{m: make(map[string]interface{}, 100)}
}

// InMemoryProxy sets the value in the map object as a cache
type InMemoryProxy struct {
	m map[string]interface{}
}

// Get checks if the cache is stored in the map object and returns true and the value if the cache is set
// It returns false if the value is not set
func (p *InMemoryProxy) Get(key string) (bool, interface{}, error) {
	val, ok := p.m[key]
	return ok, val, nil
}

// Set sets a value to the map object as a caches
func (p *InMemoryProxy) Set(key string, val interface{}) error {
	p.m[key] = val
	return nil
}
