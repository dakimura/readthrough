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
