package handlers

import (
	"crud_api/internals/model"
	"crud_api/internals/store"
	"encoding/json"
	"net/http"
)

func AddBook(w http.ResponseWriter, r *http.Request){
	var newBook model.Data
	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	model.Mu.Lock()
	model.Books = append(model.Books, newBook)
	model.Mu.Unlock()

	w.WriteHeader(http.StatusAccepted)
	store.Store()
	json.NewEncoder(w).Encode(newBook)
}