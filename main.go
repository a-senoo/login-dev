package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	})

	port := "8080"
	fmt.Println("Starting server on port " + port)
	http.ListenAndServe(":"+port, nil)
}
