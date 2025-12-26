package handlers

import (
	"crud_api/internals/model"
	"encoding/json"
	"net/http"
	"strconv"
)

func UpdateBook(w http.ResponseWriter, r *http.Request){
	id, _ := strconv.Atoi(r.PathValue("id"))
	model.Mu.Lock()
	defer model.Mu.Unlock()
	for i, book := range model.Books{
		if id == book.Id{
			if err := json.NewDecoder(r.Body).Decode(&model.Books[i]); err != nil{
				http.Error(w, "Invalid Json", http.StatusBadRequest)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(model.Books[i])
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}