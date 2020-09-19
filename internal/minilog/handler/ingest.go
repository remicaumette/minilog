package handler

import (
	"fmt"
	"github.com/remicaumette/minilog/internal/minilog/store"
	"io/ioutil"
	"net/http"
)

func HandleIngest(store *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, _ := ioutil.ReadAll(r.Body)
		fmt.Println(string(b))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}
}
