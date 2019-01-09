package jwt

import (
	"fmt"
	"strings"
	"testing"
)

var firstEncodeAnswer = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhMzZjMzA0OWIzNjI0OWEzYzlmODg5MWNiMTI3MjQzYyIsImV4cCI6MTQ0MjQzMDA1NCwibmJmIjoxNDQyNDI2NDU0LCJpYXQiOjE0NDI0MjY0NTR9.L1-1VoNu598QvY97rE7mLiSqmJInC9Yb5p-JUvYYUxo"
var secondEncodeAnswer = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJVTEVQNDhTZWtlbE5RMWowOWRrc3dub0MxRWhCY0dEViJ9.mabL6L1hUUtEGX8QLK_L6PiA5HO4b9fabcW1yhVZEB4"

func TestEncode(t *testing.T) {
	payload := Payload{Iss: "a36c3049b36249a3c9f8891cb127243c", Exp: 1442430054, Nbf: 1442426454, Iat: 1442426454}
	encoded, err := Encode(payload, "b0970f7fc9564e65xklfn48930b5d08b1")
	if err != nil {
		t.Error(fmt.Sprintf("Encode occurs error: %s", err))
	}

	if strings.Compare(encoded, firstEncodeAnswer) != 0 {
		t.Error(fmt.Sprintf("Enocde is wrong: %s, answer: %s", encoded, firstEncodeAnswer))
	}

	payload2 := Payload{Iss: "ULEP48SekelNQ1j09dkswnoC1EhBcGDV"}
	encoded, err = Encode(payload2, "6gOdOte0AYArWRUO7Zds62i3LDMVMweZ")
	if err != nil {
		t.Error(fmt.Sprintf("Encode occurs error: %s", err))
	}

	if strings.Compare(encoded, secondEncodeAnswer) != 0 {
		t.Error(fmt.Sprintf("Encode occurs error: %s", err))
	}
}