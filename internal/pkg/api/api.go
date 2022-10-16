package api

import (
	"encoding/json"
	"net/http"
)

type MsgResponse struct {
	StatusCode int    `json:"-"`
	Message    string `json:"message"`
}

type DataResponse []byte

var (
	ErrInternal = &MsgResponse{
		StatusCode: http.StatusInternalServerError,
		Message:    "internal error",
	}

	ErrNotFound = &MsgResponse{
		StatusCode: http.StatusNotFound,
		Message:    "not found",
	}

	ErrNotAllowed = &MsgResponse{
		StatusCode: http.StatusMethodNotAllowed,
		Message:    "method not allowed",
	}

	ErrUnauthorized = &MsgResponse{
		StatusCode: http.StatusUnauthorized,
		Message:    "unauthorized",
	}
)

func (m *MsgResponse) Data() (resp DataResponse) {
	resp, _ = json.Marshal(m)
	return
}

func (d DataResponse) Responder(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	_, err := w.Write(d)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (m *MsgResponse) Responder(w http.ResponseWriter) {
	m.Data().Responder(w, m.StatusCode)
}

func (m *MsgResponse) Handler(w http.ResponseWriter, r *http.Request) {
	m.Data().Responder(w, m.StatusCode)
}

func JSONResponder(w http.ResponseWriter, statusCode int, obj interface{}) {
	resp, err := json.Marshal(obj)
	if err != nil {
		ErrInternal.Responder(w)
	}

	DataResponse(resp).Responder(w, statusCode)
}
