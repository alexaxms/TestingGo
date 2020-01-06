package http_score

import (
	"fmt"
	"net/http"
)

func PlayerServer(response http.ResponseWriter, request *http.Request) {
	player := request.URL.Path[len("/players/"):]
	fmt.Fprint(response, ObtainPlayerScore(player))
}

func ObtainPlayerScore(name string) string {
	if name == "Alex" {
		return "20"
	}

	if name == "Alejandra" {
		return "35"
	}
	return ""
}
