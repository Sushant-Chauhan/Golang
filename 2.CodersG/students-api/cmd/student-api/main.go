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
	"github.com/Sushant-Chauhan/Go/Coders/students-api/internal/storage/sqlite"
)

func main() {
	fmt.Println("Welcome to students-api")

	// load config
	cfg := config.MustLoad()
	fmt.Printf("Config: %+v\n", cfg)

	// database setup
	_, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Storage initialized successfully", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))

	// setup router
	router := http.NewServeMux()

	// setup custom logger : here i will be using inbuild logger
	// router.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Welcome to Students API !!"))
	// })
	c router.HandleFunc("POST /api/students", student.New())

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
	err = server.Shutdown(ctx)
	if err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}
	slog.Info("Server Shutdown Successfully ...")
}

// Run :(with congig file in parameter) from the root (where the project starts)
// go run .\cmd\student-api\main.go -config .\config\local.yaml

// go env CGO_ENABLED  //check if CGO_ENABLED is set to 0 or 1
// set CGO_ENABLED=1
