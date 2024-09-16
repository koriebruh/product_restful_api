package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	/// SAVE HASIL DECODE
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(request)
	IfErrNotNul(err)
}

func WriteToResponseJson(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encode := json.NewEncoder(writer)
	err := encode.Encode(response)
	IfErrNotNul(err)
}
