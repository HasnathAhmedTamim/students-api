package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hasnathahmedtamim/students-api/internal/config"
	"github.com/hasnathahmedtamim/students-api/internal/http/handlers/student"
	"github.com/hasnathahmedtamim/students-api/internal/storage/sqlite"
)

func main() {

	// load config
	cfg := config.MustLoad()

	// database connection
	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	slog.Info("storage initialized successfully", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))

	// setup routes
	router := http.NewServeMux()

	// post student
	router.HandleFunc("POST /api/students", student.New(storage))
	// get student by id
	router.HandleFunc("GET /api/students/{id}", student.GetById(storage))
	// get all students
	router.HandleFunc("GET /api/students", student.GetAll(storage))
	// update student by id
	router.HandleFunc("PUT /api/students/{id}", student.UpdateById(storage))
	// delete student by id
	router.HandleFunc("DELETE /api/students/{id}", student.DeleteById(storage))

	// Setup server
	server := http.Server{
		Addr:    cfg.HTTPServer.Address,
		Handler: router,
	}

	slog.Info("Starting server...", slog.String("address", cfg.HTTPServer.Address))

	// Graceful shutdown

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	<-done

	slog.Info("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("Server gracefully stopped")

}
