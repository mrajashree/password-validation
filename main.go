package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type responseError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type password struct {
	Secret string
}

func handler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var p password
	err := decoder.Decode(&p)

	if err != nil {
		fmt.Printf("Error : %v\n", err)
		return
	}

	userPassword := p.Secret

	if len(userPassword) < 6 {
		errorMessage := responseError{"error", "Password length should be greater than or equal to 6!"}

		jsonErrorMessage, err := json.Marshal(errorMessage)
		if err != nil {
			fmt.Printf("Error : %v\n", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, string(jsonErrorMessage), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8090", nil)
}
