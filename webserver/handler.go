package webserver

import "net/http"

func Home(w http.ResponseWriter, r *http.Request) {
	resp := Response{
		Message: "It works!",
	}
	resp.Write(w)
}
