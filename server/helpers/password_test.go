package helpers

import "testing"

func TestHash(t *testing.T) {
	password := "password"
	wantHash, err := GenerateHash(password)

	if err != nil {
		t.Fatal(err)
	}

	correctPassword := password
	incorrectPassword := "password2"

	got := CompareHash(correctPassword, wantHash)
	if got != true {
		t.Errorf("Hash for same passwords should match hash :  - %s, password %s", wantHash, correctPassword)
	}

	got = CompareHash(incorrectPassword, wantHash)
	if got == true {
		t.Errorf("Hash for different passwords should not match :  hash - %s, password %s", wantHash, correctPassword)
	}

}
