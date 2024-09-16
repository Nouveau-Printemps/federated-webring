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
	UpdateEndpoint  string `json:"update_endpoint"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	resp := webserver.NewResponse(http.StatusOK, "Important information")
	data := HelloData{
		Name:            "A name",
		Version:         "1",
		ApplicationName: "Federated WebRing",
		Description:     "A description",
		UpdateEndpoint:  "/api/update",
	}
	resp.Data = data
	resp.Write(w)
}
