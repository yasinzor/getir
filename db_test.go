package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateNewMongoClient(t *testing.T) {
	mongoClient, ctx, err := CreateNewMongoClient()
	require.Nil(t, err)
	require.NotNil(t, mongoClient)
	require.NotNil(t, ctx)
}