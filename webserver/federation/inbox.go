package federation

import (
	"encoding/json"
	"github.com/Nouveau-Printemps/federated-webring/webserver"
	"io"
	"log/slog"
	"net/http"
	"strings"
)

var typeMap map[string]func(http.ResponseWriter, *Data)

func init() {
}

func InboxHandler(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		webserver.NewResponse(http.StatusInternalServerError, "An error occurred").Write(w)
		slog.Error(err.Error())
		return
	}
	var reqData Data
	err = json.Unmarshal(b, &reqData)
	if err != nil {
		webserver.NewResponse(http.StatusInternalServerError, "An error occurred").Write(w)
		slog.Error(err.Error())
		return
	}
	if strings.HasPrefix(reqData.Type, "valid/") {
		reqData.Valid(w)
		return
	}
	valid, err := reqData.Verify()
	if err != nil {
		webserver.NewResponse(http.StatusInternalServerError, "An error occurred").Write(w)
		slog.Error(err.Error())
		return
	}
	if valid {
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
	f, ok := typeMap[reqData.Type]
	if !ok {
		webserver.NewResponse(http.StatusBadRequest, "Type not found").Write(w)
	}
	f(w, &reqData)
}
