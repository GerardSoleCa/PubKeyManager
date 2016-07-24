package domain

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"strings"
)

type KeyRepository interface {
	Store(key *Key) error
	Delete(id int64) error
	GetKeys() []Key
	GetUserKeys(user string) []Key
}

type Key struct {
	Id          int64  `json:"id,omitempty"`
	User        string `json:"user"`
	Title       string `json:"title"`
	Fingerprint string `json:"fingerprint"`
	Key         string `json:"key"`
}

func (k *Key) CalculateFingerprint() {
	var fingerprint []string

	b, _ := base64.StdEncoding.DecodeString(strings.Split(k.Key, " ")[1])

	h := md5.New()
	h.Write(b)
	hash := hex.EncodeToString(h.Sum(nil))
	for i, c := range hash {
		fingerprint = append(fingerprint, string(c))
		if i != len(string(hash))-1 && i%2 == 1 {
			fingerprint = append(fingerprint, ":")
		}
	}

	k.Fingerprint = strings.Join(fingerprint, "")
}
