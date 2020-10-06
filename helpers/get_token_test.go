package helpers

import "testing"

func TestGetToken(t *testing.T) {
	res, err := GetToken(GetTokenRequest{
		Username:     "test",
		Password:     "test",
		ClientID:     "test",
		ClientSecret: "test",
	})

	if err != nil {
		t.Error(err)
	}

	if res == nil {
		t.Error("Expected response, but got nil")
	}
}
