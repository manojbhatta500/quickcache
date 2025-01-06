package main

import (
	"fmt"
	"log"
	"net/http"
	"quickcache/cache"
	"quickcache/handlers"
)

func main() {
	node1 := cache.CreateNode("apple")
	cache.CacheData.Head = &node1
	http.HandleFunc("/add", handlers.AddCache)
	http.HandleFunc("/show", handlers.ShowCache)
	fmt.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
	// cacheData.Traverse()
}
