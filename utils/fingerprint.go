package utils

import (
	"crypto/md5"
	"io"
	"fmt"
	"strings"
)

func KeyFingerprint(key string) string {
	var fingerprint []string
	h := md5.New()
	io.WriteString(h, key)
	hash := fmt.Sprintf("%x", h.Sum(nil))
	for i, c := range hash {
		fingerprint = append(fingerprint, string(c))
		if i != len(string(hash))-1 && i%2 == 1 {
			fingerprint = append(fingerprint, ":")
		}
	}
	return strings.Join(fingerprint, "")
}