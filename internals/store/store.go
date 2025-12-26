package store

import (
	"crud_api/configs"
	"crud_api/internals/model"
	"encoding/json"
	"os"
)

func Store(){
	res, _ := json.MarshalIndent(&model.Books, "", " ")
	os.WriteFile(configs.JsonPath, res, 0644)
}