package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
		return
	}

	data := models.Person{}
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
	w.WriteHeader(http.StatusOK)
	res := models.Response{Ok: "ok", Msg: "persona guardada"}
	js, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
}

func (db *database) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusBadRequest)

		res := models.Response{Ok: "false", Msg: "metodo no permitido"}

		js, err := json.Marshal(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(js)
		return
	}

	// lee todos los query params
	// validar que el numero sea positivo
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		w.Write([]byte(`
			{
				"ok" : "false"
				"message" : "El id debe ser un numero valido"
			}
		`))
		return
	}

	data := models.Person{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`
			{
				"ok" : "false"
				"message" : "La persona no tiene una estructura correcta"
			}
		`))

		return
	}

	err = db.storage.Update(id, &data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`
			{
				"ok" : "false"
				"message" : "no se pudo guardar"
			}
		`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`
			{
				"ok" : "true"
				"message" : "persona actualizada"
			}
		`))
}

func (db *database) getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)

		res := models.Response{Ok: "false", Msg: "metodo no permitido"}

		js, err := json.Marshal(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(js)
		return
	}

	resp, err := db.storage.GetAll()
	// problema al consultar la base de datos
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := models.Response{Ok: "false", Msg: "problema al consultar la bd"}
		js, err := json.Marshal(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(js)

		return
	}

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
