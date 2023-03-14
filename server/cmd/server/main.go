package main

import (
	"chat/server/internal/server"
	"net/http"
)

func main() {
	http.HandleFunc("/", server.HelloServer)
	http.ListenAndServe(":8080", nil)
}
