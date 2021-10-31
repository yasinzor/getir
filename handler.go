package main

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"log"
	"net/http"
	"time"
)
// MemoryDb decision to get or post method runs
func (s *Server) MemoryDb (w http.ResponseWriter, r *http.Request) {
	if GET == r.Method {
		GetInMemoryKey(s,w, r)
	} else if POST == r.Method {
		PostInMemoryKeyVal(s, w, r)
	}
}

// GetCountOfKey represent to getting report which is count of keys.
func (s *Server) GetCountOfKey(w http.ResponseWriter, r *http.Request) {
	var (
		requestBody  Request
		response Response
		recordBody []RecordBody
	)

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		errorResponseJSON := createErrorResponce(err)
		createHttpResponse(w, http.StatusInternalServerError, errorResponseJSON, err)
	}
	col := s.MongoClient.Database(DEFAULT_MONGO_DATABASE).Collection("records")

	startDate, err := time.Parse("2006-01-02", requestBody.StartDate)
	if err != nil {
		log.Println(err)
		errorResponseJSON := createErrorResponce(err)
		createHttpResponse(w, http.StatusInternalServerError, errorResponseJSON, err)
	}

	endDate, err := time.Parse("2006-01-02", requestBody.EndDate)
	if err != nil {
		log.Println(err)
		errorResponseJSON := createErrorResponce(err)
		createHttpResponse(w, http.StatusInternalServerError, errorResponseJSON, err)
	}
	pipeline := mongo.Pipeline{
		{{"$match", bson.D{
			{"createdAt", bson.D{{"$lte", endDate}}},
			{"createdAt", bson.D{{"$gte", startDate}}},
		}}},
		{{"$project", bson.D{
			{"key", 1},
			{"createdAt", 1},
			{"totalCount", bson.D{{"$reduce", bson.D{{"input", "$counts"}, {"initialValue", "[ ]"}, {"in", bson.D{{"$sum", "$counts"}}}}}}}}}},
		{{"$match", bson.D{
			{"totalCount", bson.D{{"$lte", requestBody.MaxCount}}},
			{"totalCount", bson.D{{"$gte", requestBody.MinCount}}},
		}}},
	}

	data, err := col.Aggregate(s.Context, pipeline)
	if err != nil {
		log.Println(err)
		errorResponseJSON := createErrorResponce(err)
		createHttpResponse(w, http.StatusInternalServerError, errorResponseJSON, err)
	}

	if err = data.All(s.Context, &recordBody); err != nil {
		log.Println(err)
		errorResponseJSON := createErrorResponce(err)
		createHttpResponse(w, http.StatusInternalServerError, errorResponseJSON, err)
	}
	response.Code = SUCCESS_CODE
	response.Message = SUCCESS_MESSAGE
	response.Records = recordBody

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		errorResponseJSON := createErrorResponce(err)
		createHttpResponse(w, http.StatusInternalServerError, errorResponseJSON, err)
	}
	createHttpResponse(w, http.StatusOK, jsonResponse, err)
}