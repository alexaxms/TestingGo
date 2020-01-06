package main

import(
	"awesomeProject/http_score"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server Running in http://localhost:5000")
	handler := http.HandlerFunc(http_score.PlayerServer)
	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}