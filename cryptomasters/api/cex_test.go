package api_test

import (
	"testing"

	"frontendmasters.com/go/crypto/api"
)

func TestApiCall(t *testing.T) {
	_, err := api.GetRate("")
	if err == nil {
		t.Error("Error was not found")
	}
}
