package webserver

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type Response struct {
	Status  uint   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type RingData struct {
	Name            string
	Host            string
	ProtocolVersion string
	ApplicationName string
	Description     string
}

var Data *RingData

func NewResponse(status uint, message string) *Response {
	return &Response{
		Status:  status,
		Message: message,
	}
}

func (r *Response) Write(w http.ResponseWriter) {
	r.write(w, false)
}

func (r *Response) write(w http.ResponseWriter, force bool) {
	if r.Status == 0 {
		r.Status = http.StatusOK
	}
	b, err := json.Marshal(r)
	if err != nil {
		if force {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			slog.Error(err.Error())
			return
		}
		generateInternalError(err).write(w, true)
		slog.Error(err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(b)
	if err != nil {
		generateInternalError(err).write(w, true)
		slog.Error(err.Error())
	}
}

func generateInternalError(err error) *Response {
	return &Response{
		Status:  http.StatusInternalServerError,
		Message: err.Error(),
		Data:    nil,
	}
}
