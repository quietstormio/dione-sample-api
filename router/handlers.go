package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Cloud struct {
	Name  string
	Storm bool
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the Home page!")
}

func storm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var c Cloud

	myCloud := `{"name": "cumulonimbus", "storm": true}`

	// look up marshalling JSON
	json.Unmarshal([]byte(myCloud), &c)

	fmt.Fprintf(w, "%+v", c)
}
