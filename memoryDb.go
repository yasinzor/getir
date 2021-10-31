package main

import (
	"encoding/json"
	"net/http"
)

// PostInMemoryKeyVal is save da key value pair in memory database
func PostInMemoryKeyVal(s *Server, w http.ResponseWriter, r *http.Request) {
	body, err := RequestBody(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	contentType := r.Header.Get(CONTENT_TYPE)
	if APPLICATION_JSON != contentType {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Request body not a JSON`))
		return
	}

	keyVal := ""
	valueVal := ""
	if key, ok := body["key"].(string); ok {
		keyVal = key
	} else {
		// Bad request
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`"key" is missing`))
		return
	}

	if value, ok := body["value"].(string); ok {
		valueVal = value
	} else {
		// Bad request
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`"value" is missing`))
		return
	}

	s.InmemoryDatabase[keyVal] = valueVal
	response := map[string]interface{}{"key": keyVal, "value": valueVal}
	resBytes, _ := json.Marshal(response)
	setJSONResponseHeader(w, http.StatusOK)
	w.Write(resBytes)
}

// GetInMemoryKey is getting data in memory database using filtering key
func GetInMemoryKey(s *Server,w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	key := values.Get("key")
	if key == "" {
		// No query param, return empty
		response := map[string]interface{}{}
		resBytes, _ := json.Marshal(response)
		setJSONResponseHeader(w, http.StatusOK)
		w.Write(resBytes)
		return
	}

	if value, ok := s.InmemoryDatabase[key]; ok {
		response := map[string]interface{}{"key": key, "value": value}
		resBytes, _ := json.Marshal(response)
		setJSONResponseHeader(w, http.StatusOK)
		w.Write(resBytes)
	}
}
