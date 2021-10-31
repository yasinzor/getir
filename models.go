package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	)
// Server represent main proxy is it contain mongoclient context memory database
type Server struct {
	// InmemoryDatabase
	InmemoryDatabase map[string]string
	// MongoClient
	MongoClient      *mongo.Client
	// Context
	Context          context.Context
}

// Response represent general responce sturct
type Response struct {
	// Code
	Code    int          `json:"code"`
	// Message
	Message string       `json:"msg"`
	// Records
	Records []RecordBody `json:"records"`
}

// RecordBody represent total count of key
type RecordBody struct {
	// Key
	Key        string    `json:"key"`
	// CreatedAt
	CreatedAt  time.Time `json:"createdAt"`
	// TotalCount
	TotalCount int32     `json:"totalCount"`
}

// Request reprsent when getting total count of key searching and filter params
type Request struct {
	// StartDate
	StartDate string `json:"startDate"`
	// EndDate
	EndDate   string `json:"endDate"`
	// MinCount
	MinCount  int32  `json:"minCount"`
	// MaxCount
	MaxCount  int32  `json:"maxCount"`
}