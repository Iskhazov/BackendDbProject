package auth

import "testing"

func TestHushPassword(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("Error hashing paswword: %v", err)
	}
	if hash == "" {
		t.Errorf("EmptyHash")
	}
	if hash == "password" {
		t.Errorf("Hash equals password")
	}
}
func TestComparePasswords(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("TestComparePasswords: Error hashing paswword: %v", err)
	}

	if !ComparePasswords(hash, []byte("password")) {
		t.Errorf("TestComparePasswords: expected password to match hash")
	}

	if ComparePasswords(hash, []byte("notpassword")) {
		t.Errorf("TestComparePasswords: expected password NOT to match hash")
	}

}
