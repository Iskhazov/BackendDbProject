package auth

import "testing"

func TestCreateJWT(t *testing.T) {
	secret := []byte("secret")

	token, err := CreateJWT(secret, 1)
	if err != nil {
		t.Errorf("TestCreateJWT: Error creating JWT: %v", err)
	}
	if token == "" {
		t.Errorf("TestCreateJWT: token is empty")
	}
}
