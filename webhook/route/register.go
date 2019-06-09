package route

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Register method to registering a account connect to bot
func Register(inputArr []string) string {

	// input must have 2 parameter
	if len(inputArr) < 4 {
		return "Parameter yang anda inputkan salah"
	}

	// call api for log in
	jsonData := map[string]string{"email": inputArr[1], "password": inputArr[2], "line_id": inputArr[3], "type": inputArr[4]}
	jsonValue, _ := json.Marshal(jsonData)

	response, _ := http.Post("http://portal.olimpiade.id/api/line-subscription", "application/json", bytes.NewBuffer(jsonValue))

	if response.StatusCode != 200 { // login failed as long as response code is not equal with 200
		return "login gagal"
	}

	return "Line anda telah berhasil terhubung"
}
