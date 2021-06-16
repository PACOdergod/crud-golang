package handler

import "net/http"

func RouterPerson(mux *http.ServeMux, storage Storage) {

	h := NewDatabase(storage)

	mux.HandleFunc("/v1/person/create", h.Create)
	mux.HandleFunc("/v1/person/update", h.update)
	mux.HandleFunc("/v1/person/get-all", h.getAll)
}
