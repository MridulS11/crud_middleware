package handlers

import (
	"crud_api/internals/model"
	"encoding/json"
	"net/http"
	"strconv"
)

func UpdateBookPart(w http.ResponseWriter, r * http.Request){
	id, err := strconv.Atoi(r.PathValue("id"))
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }
	var input model.UpdateData
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil{
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}
	model.Mu.Lock()
	defer model.Mu.Unlock()
	for i := range model.Books{
		if model.Books[i].Id == id{
			if input.Author != nil{
				model.Books[i].Author = *input.Author
			}
			if input.Title != nil{
				model.Books[i].Title = *input.Title
			}
			json.NewEncoder(w).Encode(model.Books[i])
			return
		}
	}
	http.Error(w, "Book Not Found", http.StatusNotFound)
}