package handlers

import (
	"crud_api/internals/model"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetBookById(w http.ResponseWriter, r *http.Request){
	id := r.PathValue("id")
	i, _ := strconv.Atoi(id)
	w.Header().Set("Content-Type", "application/json")
	for _, book := range model.Books{
		if book.Id == i{
			if err := json.NewEncoder(w).Encode(book); err != nil{
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}
	}
}