package http_score

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpScore(t *testing.T) {
	t.Run("return Alex's score", func(t *testing.T) {
		response := httptest.NewRecorder()
		PlayerServer(response, createNewRequest("Alex"))

		assertResponseBody(t, "20", response.Body.String())

	})

	t.Run("return Alejandra's score", func(t *testing.T) {
		response := httptest.NewRecorder()
		PlayerServer(response, createNewRequest("Alejandra"))

		assertResponseBody(t, "35", response.Body.String())
	})
}

func assertResponseBody(t *testing.T, expected string, result string) {
	if result != expected {
		t.Errorf("result '%s', expected '%s'", result, expected)
	}
}

func createNewRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}