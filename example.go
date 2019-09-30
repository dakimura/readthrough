package main

import (
	"fmt"
	"github.com/dakimura/readthrough/read"
	"github.com/dakimura/readthrough/proxy"
	"time"
)

var rt = read.Through{Proxy: proxy.NewInMemoryProxy()}

func main() {
	// 1st exec takes 5 sec
	try()

	// 2nd exec takes almost 0 sec
	try()
}

func try() {
	val, _ := rt.Get("foo",
		func() (interface{}, error) {
			// Some slow process (e.g. http access, DB access)
			time.Sleep(5 * time.Second)
			return "hello", nil
		},
	)

	data, _ := val.(string)
	fmt.Println(data)
}
