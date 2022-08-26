package main

import (
	"fmt"
	"net/http"
	"os"
)

var (
	port    = ":8080"
	version = "v1.0"
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from API %s!\n", version)
	fmt.Println("Endpoint Hit: root")
}
func handleRequests() {
	http.HandleFunc("/", root)
	http.ListenAndServe(port, nil)
}

func main() {
	if len(os.Args) > 1 {
		port = ":" + os.Args[1]
		if len(os.Args) > 2 {
			version = os.Args[2]
		}
	}
	fmt.Printf("API version: %s\n", version)
	fmt.Println("Listening on port", port)
	handleRequests()
}
