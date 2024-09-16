package federation

import (
	"encoding/json"
	"errors"
	"github.com/Nouveau-Printemps/federated-webring/webserver"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
	"strings"
)

type Data struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Origin  string `json:"origin"`
	UUID    string `json:"uuid"`
}

var (
	validTypeMap         map[string]string
	validTypeFuncMap     map[string]func(http.ResponseWriter, *Data)
	ErrValidTypeNotFound = errors.New("valid type not found")
	ErrUnknownResponse   = errors.New("unknown status code in response")
)

func init() {
	validTypeMap["federation/request"] = "valid/federation-request"
	validTypeMap["federation/response"] = "valid/federation-response"
	validTypeFuncMap["valid/federation-request"] = validFederationRequest
	validTypeFuncMap["valid/federation-response"] = validFederationResponse
}

func (d *Data) Verify() (bool, error) {
	typ, ok := validTypeMap[d.Type]
	if !ok {
		return false, ErrValidTypeNotFound
	}
	data := Data{
		Type:    typ,
		Message: d.UUID,
		Origin:  webserver.Data.Host,
		UUID:    uuid.New().String(),
	}
	mb, err := json.Marshal(data)
	if err != nil {
		return false, err
	}
	req, err := http.NewRequest(http.MethodPost, d.Origin, strings.NewReader(string(mb)))
	if err != nil {
		return false, err
	}
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	switch resp.StatusCode {
	case http.StatusOK:
		return true, nil
	case http.StatusForbidden:
		{
			slog.Warn("Request is not valid", "uuid", d.UUID)
			return false, nil
		}
	}
	return false, ErrUnknownResponse
}

func (d *Data) Valid(w http.ResponseWriter) {
	f, ok := validTypeFuncMap[d.Type]
	if !ok {
		webserver.NewResponse(http.StatusBadRequest, "Type not found").Write(w)
		return
	}
	f(w, d)
}

func validFederationRequest(w http.ResponseWriter, data *Data) {
	//
}

func validFederationResponse(w http.ResponseWriter, data *Data) {
	//
}
