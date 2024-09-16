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

func RandomSiteHandler(w http.ResponseWriter, r *http.Request) {
	resp := webserver.NewResponse(http.StatusOK, "Random website")
	resp.Data = &SiteData{
		Name:        "anhgelus Blog",
		URL:         "https://blog.anhgelus.world/",
		Description: "Blog of anhgelus",
		Type:        "blog",
	}
	resp.Write(w)
}

func SiteHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	url := q.Get("url")
	name := q.Get("name")
	if name == "" && url == "" {
		webserver.NewResponse(http.StatusBadRequest, "url and name queries not set").Write(w)
		return
	}
	if name != "anhgelus Blog" && url != "https://blog.anhgelus.world/" {
		webserver.NewResponse(http.StatusNotFound, "Not found :(").Write(w)
		return
	}
	resp := webserver.NewResponse(http.StatusOK, "Found!")
	resp.Data = &SiteData{
		Name:        "anhgelus Blog",
		URL:         "https://blog.anhgelus.world/",
		Description: "Blog of anhgelus",
		Type:        "blog",
	}
	resp.Write(w)
}
