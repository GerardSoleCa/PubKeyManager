package utils

import (
	"net/http"
	"encoding/json"
	"fmt"
	"io"
)

type HttpUtils struct{}

type ApiError struct {
	Code int `json:"-"`
	Err  string `json:"error,omitempty"`
}

func (h HttpUtils) ParseBody(body io.Reader, t interface{}) (error) {
	decoder := json.NewDecoder(body)
	return decoder.Decode(t)
}

func (h HttpUtils) Ok(rw http.ResponseWriter, body string) {
	h.send(rw, body, 200)
}

func (h HttpUtils) Created(rw http.ResponseWriter) {
	h.send(rw, "", 201)
}
func (h HttpUtils)CreatedWithBody(rw http.ResponseWriter, body interface{}) {
	b, _ := json.Marshal(body)
	h.send(rw, string(b), 201)
}

func (h HttpUtils) NoContent(rw http.ResponseWriter) {
	h.send(rw, "", 204)
}

func (h HttpUtils) BadRequest(rw http.ResponseWriter) {
	h.ErrorResponse(rw, &ApiError{Code: 400, Err: "Bad Request"})
}

func (h HttpUtils) NotFound(rw http.ResponseWriter) {
	h.ErrorResponse(rw, &ApiError{Code: 404, Err: "Not Found"})
}

func (h HttpUtils) InternalServerError(rw http.ResponseWriter) {
	h.ErrorResponse(rw, &ApiError{Code: 500, Err: "Internal Server Error"})
}

func (h HttpUtils) Unauthorized(rw http.ResponseWriter) {
	h.ErrorResponse(rw, &ApiError{Code: 401, Err: "Unauthorized"})
}

func (h HttpUtils) ErrorResponse(rw http.ResponseWriter, err *ApiError) {
	b, _ := json.Marshal(err)
	h.send(rw, string(b), err.Code)
}

func (h HttpUtils) send(rw http.ResponseWriter, body string, code int) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(code)
	if body != "" {
		fmt.Fprintln(rw, body)
	}
}