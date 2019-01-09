package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"strings"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Iss    string      	`json:"iss,omitempty"`
	Exp    int64        `json:"exp,omitempty"`
	Nbf    int64    	`json:"nbf,omitempty"`
	Iat    int64        `json:"iat,omitempty"`
}

func Hmac256(src string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(src))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func Encode(payload Payload, secret string) (string, error) {
	header := Header{Alg: "HS256", Typ: "JWT"}
	str, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	encodedHeader := Base64Encode(string(str))
	serializedPayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	sigVal := encodedHeader + "." + Base64Encode(string(serializedPayload))
	token := sigVal + "." + Hmac256(sigVal, secret)
	r := strings.NewReplacer("=", "", "+", "-", "/", "_")
	return r.Replace(token), nil
}

func Base64Encode(src string) string {
	data := []byte(src)
	return base64.StdEncoding.EncodeToString(data)
}

func Base64Decode(src string) (string, error) {
	bytes, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}