package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/space", getSpaceHandler)
	http.ListenAndServe(":8080", nil)
}
