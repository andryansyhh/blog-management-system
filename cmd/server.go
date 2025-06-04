package cmd

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()
	deps := InitDependencies(router)

	server := &http.Server{
		Addr:    ":" + deps.Config.Port,
		Handler: router,
	}

	// Channel untuk shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start server di goroutine
	go func() {
		log.Printf("HTTP server running on port %s", deps.Config.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	log.Printf("Connected to DB on %s:%s", deps.Config.DBHost, deps.Config.DBPort)
	log.Printf("Connected to Redis on %s", deps.Config.RedisAddr)

	<-stop

	log.Println("Shutting down server...")
	if err := server.Close(); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}
	log.Println("Server gracefully stopped.")
}
