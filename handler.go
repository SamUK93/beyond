package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	beyond "github.com/SamUK93/beyond/service"
)

func getSpaceHandler(w http.ResponseWriter, r *http.Request) {
	events := beyond.GetEvents()
	jsonResponse, err := json.Marshal(events)
	if err != nil {
		fmt.Fprintf(w, "Error occurred when retrieving space events")
	}
	fmt.Fprintf(w, string(jsonResponse))
}
