package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	)

// GetEnvVariableOrDefault take starting parameter values
func GetEnvVariableOrDefault(key string, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		val = defaultVal
	}
	return val
}

// RequestBody read de request body and return de map[string]interface
func RequestBody(r *http.Request) (map[string]interface{}, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		return nil, err
	}

	log.Printf("RequestBody: \n%s", body)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	m := make(map[string]interface{}, 0)
	json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func setJSONResponseHeader(w http.ResponseWriter, status int) {
	h := w.Header()
	h.Set(CONTENT_TYPE, APPLICATION_JSON)
	w.WriteHeader(status)
}

// createHttpResponse is prepare the http response
func createHttpResponse(w http.ResponseWriter, status int, jsonResponse []byte, err error) {
	if jsonResponse != nil {
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(status)
		w.Write(jsonResponse)
	} else {
		w.WriteHeader(status)
		w.Write([]byte(err.Error()))
	}
}

// createErrorResponce is prepare the http response data when proccess have a error
func createErrorResponce(err error) []byte {
	response := Response{}
	response.Code = ERROR_CODE
	response.Message = err.Error()
	jsonResponse, _ := json.Marshal(response)
	return jsonResponse
}