package main

import (
	"fmt"
	"net/http"
	"simplebuy/routers"
)

func main() {
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", 5000),
		Handler:        router,
		ReadTimeout:    60,
		WriteTimeout:   60,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}