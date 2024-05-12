package main

import (
	_ "bitoOA/docs"
	"bitoOA/internal/config"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title Backend API
// @version 1.0
// @description API Documentation

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// schemes http
func main() {

	app := InitializeApp()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.New().Service.Port),
		Handler: app.SetupRoutes(),
	}

	cfg := config.New()
	go func() {
		log.Println("server start: ", cfg.Service.Port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("shutdown server")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		panic(err)
	}

	log.Println("server exit properly")
}
