package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{"Authorization": []string{"ApiKey 1234567890123456789012345678901234567890"}}

	key, err := GetAPIKey(headers)
	if err != nil {
		t.Error("Error getting API key")
	}

	if reflect.TypeOf(key).String() != "string" {
		t.Error("API key is not a string")
	}
	if len(key) != 40 {
		t.Error("API key is not 40 characters long")
	}
}
