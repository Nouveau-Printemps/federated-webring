package api

import (
	"github.com/nouveau-printemps/federated-webring/webserver"
	"net/http"
)

type HelloData struct {
	Name            string `json:"name"`
	Version         string `json:"version"`
	ApplicationName string `json:"application_name"`
	Description     string `json:"description"`
}

type SiteData struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	resp := webserver.NewResponse(http.StatusOK, "Hello!")
	data := &HelloData{
		Name:            "A name",
		Version:         "1",
		ApplicationName: "Federated WebRing",
		Description:     "A description",
	}
	resp.Data = data
	resp.Write(w)
}

func SitesHandler(w http.ResponseWriter, r *http.Request) {
	resp := webserver.NewResponse(http.StatusOK, "Sites in the ring")
	var data []*SiteData
	data = append(data, &SiteData{
		Name:        "anhgelus Blog",
		URL:         "https://blog.anhgelus.world/",
		Description: "Blog of anhgelus",
		Type:        "blog",
	})
	resp.Data = &data
	resp.Write(w)
}
