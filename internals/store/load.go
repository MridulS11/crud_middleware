package store

import (
	"crud_api/configs"
	"crud_api/internals/model"
	"encoding/json"
	"log"
	"os"
)

func Load(){

	if _, err := os.Stat(configs.JsonPath); os.IsNotExist(err){
		log.Println("File Not Creating, Initializing One!")
		model.Books = []model.Data{}
		return
	}
	f, _ := os.ReadFile(configs.JsonPath)
	json.Unmarshal(f, &model.Books)
	
}