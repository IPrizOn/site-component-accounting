package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HomePage)

	fmt.Println("Server is listening...")

	http.ListenAndServe(":8080", nil)
}
func HomePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
