package handler

import (
	"fmt"
	"net/http"
)

// Store represents a data store
type Store interface {
	Fetch() string
	Cancle()
}

// NewHandler returns a http handler
func NewHandler(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		data := make(chan string)
		go func() {
			data <- store.Fetch()
		}()
		select {
		case d := <-data:
			fmt.Fprintf(w, d)
		case <-ctx.Done():
			store.Cancle()
		}

	}
}
