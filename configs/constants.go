package configs

import "os"

const ErrCode = "Error Encountered"
const JsonPath = "./data.json"
const Key = "golem"

func Getkey() string{
	key := os.Getenv("X_API_KEY")
	if key == ""{
		return key
	}
	return key
}