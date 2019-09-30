# readthrough

## Overview
This library is to implement read-through cache quickly in your application.


In order to improve the application performance, we sometimes implement caching.
But it takes time because we need to implement a cache accessor,
integrate it to your business logic, unit tests and integration tests of it.
This library is to get a quick win to improve your application performance without changing your application code much, by a simple read-through cache.

Also you can implement your own caching (e.g. redis or memcache) to customize this library.

## Installation
```go
import "github.com/dakimura/readthrough"
```

## Usage
The usage of this library is to just wrap a slow function that you want to improve performance
like as follows:
```go
rt := readthrough.Through{Proxy: readthrough.NewInMemoryProxy()}

value, _ := rt.Get("cacheKey", someSlowFunction)
```

When `someSlowFunction` is called with the `cacheKey`, the returned values are cached with `cacheKey`. So the cached value will be returned immediately when the function is called again.

Note: We assume that your application has a slow function which returns `(interface{}, error)` .
```go
func someSlowFunction() (interface{}, error) {
	// example
	time.Sleep(3 * time.Second)
	return "hoge", nil
}
```
If the returned values are not `(interface{}, error)` 
please wrap the function and make another function which returns `(interface{}, error)` .

## Example Code

```go
package main

import (
	"fmt"
	"github.com/dakimura/readthrough"
	"time"
)

var rt = readthrough.Through{Proxy: readthrough.NewInMemoryProxy()}

func main() {
	fmt.Println("before try... current time: " + time.Now().String())

	cacheKey := "hello"
	// 1st try takes 3 seconds
	value, _ := rt.Get(cacheKey, someSlowFunction)
	fmt.Println(value.(string))

	// 2nd try takes almost 0 second
	value, _ = rt.Get(cacheKey, someSlowFunction)
	fmt.Println(value.(string))
}

func someSlowFunction() (interface{}, error) {
	time.Sleep(3 * time.Second)
	return "current time: " + time.Now().String(), nil
}
```

Output:
```go
before try... current time: 2019-09-30 19:02:02.941276 +0900 JST m=+0.000198594
current time: 2019-09-30 19:02:05.943923 +0900 JST m=+3.002818374
current time: 2019-09-30 19:02:05.943923 +0900 JST m=+3.002818374
```

## Original Proxy
In the examples above, `proxy.NewInMemoryProxy()` is used for the proxy storage.
If you want to use another storage for the cache, please implement your own `proxy.Proxy` implementation.

proxy/proxy.go
```go
// Proxy behaves as a read-through cache for an origin (DB, API server, etc)
type Proxy interface {
	Get(key string) (bool, interface{}, error)
	Set(key string, val interface{}) error
}
```