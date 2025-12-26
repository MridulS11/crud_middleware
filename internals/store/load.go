package store

import (
	"crud_api/configs"
	"crud_api/internals/model"
	"encoding/json"
	"os"
)

func Load(){
	f, _ := os.ReadFile(configs.JsonPath)
	json.Unmarshal(f, &model.Books)
}