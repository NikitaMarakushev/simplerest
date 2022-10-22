package apiserver

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIServer_Start(t *testing.T) {
	server := New(NewConfig())
	record := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	server.handleHello().ServeHTTP(record, request)
	assert.Equal(t, record.Body.String(), "Hello")
}
