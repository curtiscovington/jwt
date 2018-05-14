package jwt

import (
	"io/ioutil"
	"testing"
)

func TestEncode(t *testing.T) {
	type payload struct {
		Test string `json:"test"`
	}
	key, _ := ioutil.ReadFile("private.pem")
	_, err := Encode(payload{Test: "Hi"}, key)
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestDecode(t *testing.T) {
	type payload struct {
		Test string `json:"test"`
	}
	key, _ := ioutil.ReadFile("private.pem")
	str, _ := Encode(payload{Test: "TEST"}, key)

	var p payload
	err := Decode(str, &p, key)
	if err != nil {
		t.Errorf("%v", err)
	}

	if p.Test != "TEST" {
		t.Errorf("Decoded payload does not match")
	}
}
func TestVerify(t *testing.T) {
	type payload struct {
		Test string `json:"test"`
	}
	key, _ := ioutil.ReadFile("private.pem")
	str, err := Encode(payload{Test: "Hi"}, key)
	if err != nil {
		t.Errorf("%v", err)
	}
	if !Verify(str, key) {
		t.Errorf("Verify failed.")
	}
}
