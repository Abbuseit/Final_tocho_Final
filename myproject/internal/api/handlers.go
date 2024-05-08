package api

import (
	"encoding/json"
	"net/http"
)

func orderHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "Order created"})
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
	}
}

func menuHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		menu := []string{"Item 1", "Item 2", "Item 3"}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(menu)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
	}
}
