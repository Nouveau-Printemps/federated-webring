package api

import (
	"github.com/Nouveau-Printemps/federated-webring/data"
	"github.com/Nouveau-Printemps/federated-webring/webserver"
	"math/rand/v2"
	"net/http"
	"strings"
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

var siteData []*SiteData

func init() {
	for _, w := range data.Websites {
		var types []string
		for _, t := range w.Type {
			types = append(types, strings.ToLower(t.Name))
		}
		siteData = append(siteData, &SiteData{
			Name:        w.Name,
			URL:         w.URL,
			Description: w.Description,
			Type:        strings.Join(types, " "),
		})
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	resp := webserver.NewResponse(http.StatusOK, "Hello!")
	resp.Data = &HelloData{
		Name:            "A name",
		Version:         "1",
		ApplicationName: "Federated WebRing",
		Description:     "A description",
	}
	resp.Write(w)
}

func SitesHandler(w http.ResponseWriter, r *http.Request) {
	resp := webserver.NewResponse(http.StatusOK, "Sites in the ring")
	resp.Data = &siteData
	resp.Write(w)
}

func RandomSiteHandler(w http.ResponseWriter, r *http.Request) {
	resp := webserver.NewResponse(http.StatusOK, "Random website")
	resp.Data = siteData[rand.N(len(siteData))]
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
	for _, s := range siteData {
		if s.Name == name || s.URL == url {
			resp := webserver.NewResponse(http.StatusOK, "Found!")
			resp.Data = s
			resp.Write(w)
			return
		}
	}
	webserver.NewResponse(http.StatusNotFound, "Not found :(").Write(w)
}
