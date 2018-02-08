// Package jwt provides a simple way to verify a jwt, encode a payload to a jwt, and decode a payload from a jwt
package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

// Verify a jwt token string
func Verify(token string, secret []byte) bool {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return false
	}
	hash := hmac.New(sha256.New, secret)
	hash.Write([]byte(parts[0] + "." + parts[1]))
	sigHash := hash.Sum(nil)

	sig := base64.StdEncoding.EncodeToString(sigHash)
	fmt.Println(sig)
	fmt.Println(parts[2])
	if strings.Compare(sig, parts[2]) != 0 {
		return false
	}
	return true
}

// Encode a payload into a jwt, returns the jwt
func Encode(payload interface{}, secret []byte) (string, error) {
	type Header struct {
		Alg  string `json:"alg"`
		Type string `json:"type"`
	}
	header := Header{
		Alg:  "HS256",
		Type: "JWT",
	}
	h, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	headerStr := base64.StdEncoding.EncodeToString(h)
	p, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	payloadStr := base64.StdEncoding.EncodeToString(p)

	str := headerStr + "." + payloadStr

	hash := hmac.New(sha256.New, secret)
	hash.Write([]byte(str))
	str = str + "." + base64.StdEncoding.EncodeToString(hash.Sum(nil))
	return str, nil
}

// Decode a payload from a jwt, assumes Verify has already been called
func Decode(t string, payload interface{}, secret []byte) error {
	p, err := base64.StdEncoding.DecodeString(strings.Split(t, ".")[1])
	if err != nil {
		return err
	}
	fmt.Println(string(p))
	err = json.Unmarshal(p, payload)
	return err
}
