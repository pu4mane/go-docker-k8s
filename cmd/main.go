package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/pu4mane/go-docker-k8s-demo/pkg/config"
	httpserver "github.com/pu4mane/go-docker-k8s-demo/pkg/server/http"
	"github.com/pu4mane/go-docker-k8s-demo/pkg/version"
)

func main() {
	cfg, err := config.MustLoad()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf(
		"Starting the service...\ncommit: %s, build time: %s, release: %s",
		version.Commit, version.BuildTime, version.Release,
	)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	srv := httpserver.NewServer(cfg, version.BuildTime, version.Commit, version.Release)

	shutdown := make(chan struct{}, 1)
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			shutdown <- struct{}{}
			log.Printf("%v", err)
		}
	}()
	log.Print("The service is ready to listen and serve.")

	select {
	case killSignal := <-interrupt:
		switch killSignal {
		case os.Interrupt:
			log.Print("Got SIGINT...")
		case syscall.SIGTERM:
			log.Print("Got SIGTERM...")
		}
	case <-shutdown:
		log.Printf("Got an error...")
	}

	log.Print("The service is shutting down...")
	srv.Shutdown(context.Background())
	log.Print("Done")
}
