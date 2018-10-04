package main

import (
	"net/http"
	"sync"

	"github.com/derekpedersen/skatepark-api-go/appcfg"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel) // TODO: make this a flag

	var wg sync.WaitGroup

	skateparkRouter, _ := appcfg.NewSkateparkAPIRouter()
	wg.Add(1)
	go func() {
		defer wg.Done()
		http.ListenAndServe(":8080", skateparkRouter)
	}()

	aliveRouter, _ := appcfg.NewHealthAPIRouter()
	wg.Add(1)
	go func() {
		defer wg.Done()
		http.ListenAndServe(":8000", aliveRouter)
	}()

	wg.Wait()
}
