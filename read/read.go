package read

//var rt *read.ReadThrough

type Proxy interface {
	Get(key string) (bool, interface{}, error)
	Set(key string, val interface{}) error
}

func NewInMemoryProxy(size int) Proxy {
	return &ThroughInMemory{m: make(map[string]interface{}, size)}
}

type ThroughInMemory struct {
	m map[string]interface{}
}

func (tim *ThroughInMemory) Get(key string) (bool, interface{}, error) {
	val, ok := tim.m[key]
	return ok, val, nil
}

func (tim *ThroughInMemory) Set(key string, val interface{}) error {
	tim.m[key] = val
	return nil
}

type Through struct {
	Proxy Proxy
}

func (rt *Through) Execute(key string, req func() (interface{}, error)) (interface{}, error) {
	ok, val, err := rt.Proxy.Get(key)
	if ok {
		return val, err
	}

	val, err = req()
	if err != nil {
		return val, err
	}

	err = rt.Proxy.Set(key, val)
	if err != nil {
		return nil, err
	}

	return val, err
}
