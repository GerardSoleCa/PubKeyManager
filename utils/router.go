package utils

import (
	"io"
	"encoding/json"
)

func ParseBody(body io.Reader, t interface{}) (error) {
	decoder := json.NewDecoder(body)
	return decoder.Decode(t)
}