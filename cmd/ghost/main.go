package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/ghost/pkg/handler"
)

func main() {
	handlers := handler.Routes()
	fmt.Println("this is the fucking starting port")
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: handlers,
	}

	var wg sync.WaitGroup

	wg.Add(1)
	// go routing
	go func() {
		log.Printf("Listening on %s\n", httpServer.Addr)

		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "Error listening and serveing", err)
		}
	}()

	wg.Wait()
}
