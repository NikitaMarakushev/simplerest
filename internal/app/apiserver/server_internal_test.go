package apiserver

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"simplerest/internal/app/store/teststore"
	"testing"
)

func TestServer_ServerHTTPS(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/users", nil)
	s := newServer(teststore.New())
	s.ServerHTTP(rec, req)
	assert.NotEqual(t, rec.Code, http.StatusOK)
}
