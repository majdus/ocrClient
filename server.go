package main

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

func InitServer(router *mux.Router) (*http.Server)  {

	srv := &http.Server{
		Addr:         ":8000",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler: router,
	}

	return  srv
}

func RunServer(srv *http.Server)  {
	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
			os.Exit(-1)
		}
	}()

	WaitForCtrlC()
	ctx := context.Background()
	go func() {
		srv.Shutdown(ctx)
	}()

	<-ctx.Done()
	os.Exit(0)
}

func WaitForCtrlC() {
	var ctrlCWaiter sync.WaitGroup
	ctrlCWaiter.Add(1)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		ctrlCWaiter.Done()
	}()

	ctrlCWaiter.Wait()
}
