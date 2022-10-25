package apiserver

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"simplerest/internal/app/store/teststore"
	"testing"
)

func TestServer_ServerHTTPS(t *testing.T) {
	s := newServer(teststore.New())
	testCases := []struct {
		name        string
		payload     interface{}
		exectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    "user@example.org",
				"password": "password",
			},
			exectedCode: http.StatusCreated,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/users", b)
			s.ServerHTTP(rec, req)
			assert.Equal(t, tc.exectedCode, rec.Code)
		})
	}
}
