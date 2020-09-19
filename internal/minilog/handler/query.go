package handler

import (
	"encoding/json"
	"github.com/remicaumette/minilog/internal/minilog/store"
	"net/http"
)

func HandleQuery(store *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		records, err := store.Query(r.URL.Query().Get("q"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(records)
	}
}
