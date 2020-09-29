package main

import (
	"github.com/chenminhua/rockcache/cache"
	"github.com/chenminhua/rockcache/http"
	"github.com/chenminhua/rockcache/tcp"
)

func main() {
	ca := cache.New("rocksdb")
	ca2 := cache.New("inmemory")
	go tcp.New(ca).Listen()
	http.New(ca2).Listen()
}
