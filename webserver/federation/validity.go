package federation

import (
	"encoding/json"
	"errors"
	"github.com/Nouveau-Printemps/federated-webring/data"
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
	validTypeFuncMap     map[string]string
	ErrValidTypeNotFound = errors.New("valid type not found")
	ErrUnknownResponse   = errors.New("unknown status code in response")
)

func init() {
	validTypeMap["federation/request"] = "valid/federation-request"
	validTypeMap["federation/response"] = "valid/federation-response"
	validTypeMap["federation/break"] = "valid/federation-break"
	validTypeFuncMap["valid/federation-request"] = data.FederationRequestType
	validTypeFuncMap["valid/federation-response"] = data.FederationResponseType
	validTypeFuncMap["valid/federation-break"] = data.FederationBreakType
}

func (d *Data) Verify() (bool, error) {
	typ, ok := validTypeMap[d.Type]
	if !ok {
		return false, ErrValidTypeNotFound
	}
	da := Data{
		Type:    typ,
		Message: d.UUID,
		Origin:  webserver.Data.Host,
		UUID:    uuid.New().String(),
	}
	mb, err := json.Marshal(da)
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
			slog.Warn("Request is not valid", "uuid", d.UUID, "origin", d.Origin)
			return false, nil
		}
	}
	return false, ErrUnknownResponse
}

func (d *Data) Valid(w http.ResponseWriter) {
	t, ok := validTypeFuncMap[d.Type]
	if !ok {
		webserver.NewResponse(http.StatusBadRequest, "Type not found").Write(w)
		return
	}
	val, err := data.GetFederationSent(t, d.Origin)
	if err != nil {
		if !errors.Is(err, data.ErrNotFound) {
			webserver.NewResponse(http.StatusInternalServerError, "An error occurred").Write(w)
			slog.Error(err.Error())
			return
		}
		slog.Warn("Not valid request", "origin", d.Origin, "uuid", d.UUID)
		webserver.NewResponse(http.StatusForbidden, "Not valid").Write(w)
		return
	}
	if val == d.Message {
		webserver.NewResponse(http.StatusOK, "Valid").Write(w)
		return
	}
	slog.Warn("Not valid request", "origin", d.Origin, "uuid", d.UUID)
	webserver.NewResponse(http.StatusForbidden, "Not valid").Write(w)
}
