package handlers

import (
	"crud_api/internals/model"
	"encoding/json"
	"net/http"
)

func GetBook(w http.ResponseWriter, r *http.Request){
	if err := json.NewEncoder(w).Encode(model.Books); err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}