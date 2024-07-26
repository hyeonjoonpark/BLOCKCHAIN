package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("public")))
	fmt.Println("Server is running on port 8081")
	http.ListenAndServe(":8081", nil)
}
