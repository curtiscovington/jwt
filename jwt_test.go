package jwt

import (
	"testing"
)

func TestEncode(t *testing.T) {
	type payload struct {
		Test string `json:"test"`
	}

	_, err := Encode(payload{Test: "Hi"}, []byte("password"))
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestDecode(t *testing.T) {
	type payload struct {
		Test string `json:"test"`
	}

	str, _ := Encode(payload{Test: "TEST"}, []byte("password"))

	var p payload
	err := Decode(str, &p, []byte("password"))
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

	str, err := Encode(payload{Test: "Hi"}, []byte("password"))
	if err != nil {
		t.Errorf("%v", err)
	}
	if !Verify(str, []byte("password")) {
		t.Errorf("Verify failed.")
	}
}
