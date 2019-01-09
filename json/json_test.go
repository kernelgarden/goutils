package json

import (
	"strings"
	"testing"
)

func TestFromJson(t *testing.T) {
	rawJson := `
  {
    "name": "Molecule Man",
    "age": 29,
    "secretIdentity": "Dan Jukes",
    "powers": [
      "Radiation resistance",
      "Turning tiny",
      "Radiation blast"
    ]
  }
`

	json, err := FromJson(rawJson)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if json["name"] != "Molecule Man" {
		t.Fail()
	}

	age := int(json["age"].(float64))
	if age != 29 {
		t.Fail()
	}

	if ret := strings.Compare(json["secretIdentity"].(string), "Dan Jukes"); ret != 0 {
		t.Log(ret)
		t.Fail()
	}

	powers := json["powers"].([]interface{})

	if len(powers) != 3 {
		t.Fail()
	}

	first := powers[0].(string)
	if strings.Compare(first, "Radiation resistance") != 0 {
		t.Fail()
	}
}