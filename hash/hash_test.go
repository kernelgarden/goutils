package hash

import "testing"

func TestEncryptPassword(t *testing.T) {
	pw := "password"
	hash, err := EncryptPassword(pw)
	if err != nil {
		t.Error("HashPassword occurs error!")
	}

	if pw == hash {
		t.Error("Not changed before hash and after hash!")
	}
}

func TestValidatePassword(t *testing.T) {
	pw := "password"
	hash, _ := EncryptPassword(pw)

	isValiated := ValidatePassword(hash, pw)
	if isValiated != true {
		t.Error("ValidatePassword has some error!")
	}

	isValiated = ValidatePassword(hash, pw + "2")
	if isValiated == true {
		t.Error("ValidatePasword has some error!")
	}
}