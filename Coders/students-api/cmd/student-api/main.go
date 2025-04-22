package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Sushant-Chauhan/Go/Coders/students-api/internal/config"
)

func main() {
	fmt.Println("Welcome to students-api")

	// load config
	cfg := config.MustLoad()
	fmt.Printf("Config: %+v\n", cfg)

	// setup custom logger : here i will be using inbuild logger
	// setup database

	// setup router
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Students API !!"))
	})

	// setup server
	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	slog.Info("starting server", slog.String("address = ", cfg.Address))
	// fmt.Printf("Server started on %s", cfg.Address)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM) //READ ABOU THESE IMP.

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	<-done

	slog.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Read about Context
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	} else {
		slog.Info("server shutdown successfully")
	}
}

// Run :(with congig file in parameter)
// go run .\cmd\students-api\main.go -config .\config\local.yaml
