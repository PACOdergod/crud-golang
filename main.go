package main

import (
	"log"
	"net/http"

	"github.com/PACOdergod/crud-golang/handler"
	"github.com/PACOdergod/crud-golang/storage"
)

func main() {
	store := storage.NewMemory()
	mux := http.NewServeMux()

	handler.RouterPerson(mux, &store)

	log.Println("servidor iniciado en localhost:8000")

	// parece que no puede usar la direccion 8080
	// no se que la esta ocupando
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Println(err)
	}
}
