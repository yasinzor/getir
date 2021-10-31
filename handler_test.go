package main

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)
/*
func TestServer_GetCountOfKey(t *testing.T) {
	body := Request{
		StartDate: "2001-01-01",
		EndDate:   "2021-01-01",
		MinCount:  8000,
		MaxCount:  8200,
	}
	bodyByte, _ := json.Marshal(body)
	req, err := http.NewRequest("GET", "/count-of-key", bytes.NewBuffer(bodyByte))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	ctx := req.Context()
	var mt *mongo.Database
	// mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	//mt.Client.Database(DEFAULT_MONGO_DATABASE,options.Database()).CreateCollection(ctx,"records")
	//defer mt.Close()
	server := Server{
		InmemoryDatabase: make(map[string]string, 0),
		MongoClient:      mt,
		Context:          ctx,
	}

	server.GetCountOfKey(rr,req)
	require.Nil(t, err)
	require.Equal(t, 200, rr.Code)
}*/

func TestServer_MemoryDb(t *testing.T) {

	body := make(map[string]interface{}, 0)
	body["key"] = "test-key"
	body["value"] = "test-value"
	bodyByte, _ := json.Marshal(body)
	reader := strings.NewReader(string(bodyByte))
	req, err := http.NewRequest("POST", "/in-memory", reader)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	ctx := req.Context()
	var mongoClient *mongo.Client
	server := Server{
		InmemoryDatabase: make(map[string]string, 0),
		MongoClient:      mongoClient,
		Context:          ctx,
	}
	handler := http.HandlerFunc(server.MemoryDb)
	handler.ServeHTTP(rr, req)
	require.Nil(t, err)
	require.Equal(t, 200, rr.Code)

	req, err = http.NewRequest("GET", "/in-memory?key=test-key", nil)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	handler = http.HandlerFunc(server.MemoryDb)
	handler.ServeHTTP(rr, req)
	require.Nil(t, err)
	require.Equal(t, 200, rr.Code)


}
