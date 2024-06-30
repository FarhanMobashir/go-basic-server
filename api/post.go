package api

import (
	"encoding/json"
	"go-server/data"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var exhibition data.Exhibition
		err := json.NewDecoder(r.Body).Decode(&exhibition)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		data.Add(exhibition)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Ok"))

	} else {
		http.Error(w, "Unsupported Method", http.StatusMethodNotAllowed)
	}
}