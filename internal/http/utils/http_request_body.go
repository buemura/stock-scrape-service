package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetRequestBody(w http.ResponseWriter, r *http.Request, requestBody any) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		HandleErrorResponse(w, http.StatusInternalServerError, "Error reading request body")
		return
	}

	defer r.Body.Close()

	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		HandleErrorResponse(w, http.StatusInternalServerError, "Error parsing request body")
		return
	}
}