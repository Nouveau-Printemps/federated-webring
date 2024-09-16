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
		slog.Error(err.Error())
		webserver.NewResponse(http.StatusInternalServerError, "An error occurred").Write(w)
		return
	}
	var reqData Data
	err = json.Unmarshal(b, &reqData)
	if err != nil {
		slog.Error(err.Error())
		webserver.NewResponse(http.StatusInternalServerError, "An error occurred").Write(w)
		return
	}
	if strings.HasPrefix(reqData.Type, "valid/") {
		reqData.Valid(w)
		return
	}
	valid, err := reqData.Verify()
	if err != nil {
		slog.Error(err.Error())
		webserver.NewResponse(http.StatusInternalServerError, "An error occurred").Write(w)
		return
	}
	if !valid {
		webserver.NewResponse(http.StatusForbidden, "Not valid").Write(w)
		return
	}
	f, ok := typeMap[reqData.Type]
	if !ok {
		webserver.NewResponse(http.StatusBadRequest, "Type not found").Write(w)
		return
	}
	f(w, &reqData)
}
