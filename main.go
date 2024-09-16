package main

import (
	"context"
	"errors"
	"github.com/gorilla/mux"
	"github.com/nouveau-printemps/federated-webring/webserver"
	"github.com/nouveau-printemps/federated-webring/webserver/api"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", webserver.Home)
	r.HandleFunc("/api/hello", api.HelloHandler)
	r.HandleFunc("/api/websites", api.SitesHandler)
	r.HandleFunc("/api/website", api.SiteHandler)
	r.HandleFunc("/api/website-random", api.RandomSiteHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	slog.Info("Starting...")
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error(err.Error())
		}
	}()

	slog.Info("Started")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		panic(err)
	}
	slog.Info("Shutting down")
	os.Exit(0)
}
