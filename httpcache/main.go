package main

import (
	"github.com/chenminhua/httpcache/cache"
	"github.com/chenminhua/httpcache/http"
	"github.com/chenminhua/httpcache/tcp"
)

func main() {
	ca := cache.New("inmemory")
	go tcp.New(ca).Listen()
	http.New(ca).Listen()
}
