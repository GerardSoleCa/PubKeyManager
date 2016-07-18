package responses

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type ApiError struct {
	Code int `json:"-"`
	Err  string `json:"error,omitempty"`
}

func Ok(rw http.ResponseWriter, body string) {
	send(rw, body, 200)
}

func Created(rw http.ResponseWriter) {
	send(rw, "", 201)
}
func CreatedWithBody(rw http.ResponseWriter, body interface{}) {
	b, _ := json.Marshal(body)
	send(rw, string(b), 201)
}

func NoContent(rw http.ResponseWriter) {
	send(rw, "", 204)
}

func BadRequest(rw http.ResponseWriter) {
	ErrorResponse(rw, &ApiError{Code: 400, Err: "Bad Request"})
}

func NotFound(rw http.ResponseWriter) {
	ErrorResponse(rw, &ApiError{Code: 404, Err: "Not Found"})
}

func InternalServerError(rw http.ResponseWriter) {
	ErrorResponse(rw, &ApiError{Code: 500, Err: "Internal Server Error"})
}

func ErrorResponse(rw http.ResponseWriter, err *ApiError) {
	b, _ := json.Marshal(err)
	send(rw, string(b), err.Code)
}

func send(rw http.ResponseWriter, body string, code int) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(code)
	if body != "" {
		fmt.Fprintln(rw, body)
	}
}