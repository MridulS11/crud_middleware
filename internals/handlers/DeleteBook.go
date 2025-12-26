package handlers

import (
	"crud_api/internals/model"
	"net/http"
	"slices"
	"strconv"
)

func DeleteBook(w http.ResponseWriter, r *http.Request){
	id, _ := strconv.Atoi(r.PathValue("id"))
	found := false
	for i, book := range model.Books{
		if book.Id == id{
			model.Books = slices.Delete(model.Books, i, i+1)
			found = true
			break
		}
	}

	// newBooks := []model.Data{}
	// for _, b := range model.Books {
	// 	if b.Id != id {
	// 		newBooks = append(newBooks, b)
	// 	}
	// }
	// model.Books = newBooks

	if !found{
		http.Error(w, "Value Not Found", http.StatusNotFound)
	}

	w.WriteHeader(http.StatusNoContent)
}