package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PACOdergod/crud-golang/models"
)

type database struct {
	storage Storage
}

func NewDatabase(storage Storage) database {
	return database{storage}
}

func (db *database) PrintDatabase() {
	fmt.Println(db.storage.GetAll())
}

// handler create
func (db *database) Create(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("El metodo fue: %v\n", r.Method)

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)

		res := models.Response{Ok: "false", Msg: "metodo no permitido"}

		js, err := json.Marshal(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(js)

		// json.NewEncoder(w).Encode(res)

		return
	}

	data := models.Person{}

	// fmt.Printf("%+v", r.Body)

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)

		res := models.Response{Ok: "false", Msg: "la persona no tiene una estructura correcta"}

		js, err := json.Marshal(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(js)

		return
	}

	fmt.Println(data)

	err = db.storage.Create(&data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		w.Write([]byte(`
			{
				"ok" : "false"
				"message" : "Hubo un problema al guardar la persona"
			}
		`))

		return
	}

	// w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusCreated)

	res := models.Response{Ok: "ok", Msg: "persona guardada"}

	js, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)

	db.PrintDatabase()

}
