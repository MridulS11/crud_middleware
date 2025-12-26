package handlers

import (
	"crud_api/configs"
	"crud_api/internals/model"
	"encoding/json"
	"net/http"
	"strings"
)

func GetBook(w http.ResponseWriter, r *http.Request){

	query := r.URL.Query().Get("author")

	if query == ""{
		err := json.NewEncoder(w).Encode(model.Books)
		if err != nil{
			http.Error(w, configs.ErrCode, http.StatusBadRequest)
			return
		}
		return
	}

	filter := []model.Data{}
	for _, book := range model.Books{
		if strings.Contains(strings.ToLower(book.Author), strings.ToLower(query)){	//strings.EqualFold(book.Author, query)
			filter = append(filter, book)
		}
	}

	json.NewEncoder(w).Encode(filter)
	
}