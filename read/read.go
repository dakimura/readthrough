package read

import "github.com/dakimura/readthrough/proxy"

// Through is the
type Through struct {
	Proxy proxy.Proxy
}

// Get reads a value through the proxy and set the cache
func (rt *Through) Get(key string, req func() (interface{}, error)) (interface{}, error) {
	// Get Get Cache from Proxy
	ok, val, err := rt.Proxy.Get(key)

	// return the cache if found
	if ok {
		return val, err
	}

	// Get from origin
	val, err = req()
	if err != nil {
		return val, err
	}

	// Set the value from origin to the proxy cache
	err = rt.Proxy.Set(key, val)
	if err != nil {
		return nil, err
	}

	// return the value got from origin
	return val, err
}
