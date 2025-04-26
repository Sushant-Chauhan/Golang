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
	"github.com/Sushant-Chauhan/Go/Coders/students-api/internal/http/handlers/student"
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
	// router.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Welcome to Students API !!"))
	// })
	router.HandleFunc("POST /api/students", student.New())

	// setup server
	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	slog.Info("starting server", slog.String("address = ", cfg.Address))
	// fmt.Printf("Server started on %s", cfg.Address)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM) //more to READ ABOUt triggering of these.
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("failed to start server: %v", err)
		}
	}()
	<-done

	slog.Info("shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Read about Context
	defer cancel()

	//graceful shutdown
	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}
	slog.Info("Server Shutdown Successfully ...")
}

// Run :(with congig file in parameter) from the root (where the project starts)
// go run .\cmd\student-api\main.go -config .\config\local.yaml
