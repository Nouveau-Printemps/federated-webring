package main

import (
	"context"
	_ "embed"
	"errors"
	"github.com/Nouveau-Printemps/federated-webring/config"
	"github.com/Nouveau-Printemps/federated-webring/data"
	"github.com/Nouveau-Printemps/federated-webring/webserver"
	"github.com/Nouveau-Printemps/federated-webring/webserver/api"
	"github.com/gorilla/mux"
	"github.com/pelletier/go-toml/v2"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"
)

//go:embed default.toml
var defaultConfig []byte

func main() {
	b, err := os.ReadFile("config.toml")
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			panic(err)
		}
		err = os.WriteFile("config.toml", defaultConfig, 0600)
		if err != nil {
			panic(err)
		}
		os.Exit(0)
	}

	var cfg config.Config
	err = toml.Unmarshal(b, &cfg)
	if err != nil {
		panic(err)
	}

	data.Init(cfg.DatabaseCredentials)

	webserver.Data = &webserver.RingData{
		Name:            cfg.Name,
		ProtocolVersion: "1",
		ApplicationName: "Federated WebRing",
		Description:     cfg.Description,
	}

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
	err = srv.Shutdown(ctx)
	if err != nil {
		panic(err)
	}
	slog.Info("Shutting down")
	os.Exit(0)
}
