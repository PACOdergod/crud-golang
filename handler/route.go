package handler

import "net/http"

func RouterPerson(mux *http.ServeMux, storage Storage) {

	h := NewDatabase(storage)

	mux.HandleFunc("/v1/person/create", h.Create)
}
