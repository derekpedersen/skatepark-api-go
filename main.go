package main

import (
	"net/http"
	"sync"

	"github.com/derekpedersen/skatepark-api-go/appcfg"
	"github.com/jeanphorn/log4go"
)

func main() {
	// configure logging
	log4go.LoadConfiguration("./log4go.json")

	var wg sync.WaitGroup

	skateparkRouter, _ := appcfg.NewSkateparkAPIRouter()
	wg.Add(1)
	go func() {
		defer wg.Done()
		http.ListenAndServe(":8080", skateparkRouter)
	}()

	aliveRouter, _ := appcfg.NewAliveAPIRouter()
	wg.Add(1)
	go func() {
		defer wg.Done()
		http.ListenAndServe(":8000", aliveRouter)
	}()

	wg.Wait()
}
