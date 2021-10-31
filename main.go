package main
import (
	"fmt"
	"log"
	"net/http"
	)

func main(){
	mongoClient, ctx, err := CreateNewMongoClient()
	if err != nil {
		log.Fatal(err)
	}
	defer mongoClient.Disconnect(ctx)

	server := Server{
		InmemoryDatabase: make(map[string]string, 0),
		MongoClient:      mongoClient,
		Context:          ctx,
	}

	//// swagger:route GET /count-of-key
	//// .
	//// responses:
	//// Response  200: RecordBody
	//// This text will appear as description of your response body.
	http.HandleFunc("/count-of-key", func(w http.ResponseWriter, r *http.Request) {
		server.GetCountOfKey(w, r)
	})

	//// route GET /in-memory
	//// this take a url filter. ?key=[param]
	//// responses:
	//// Response  200: map[string]interface
	//// this return key value in memory database.
	http.HandleFunc("/in-memory", func(w http.ResponseWriter, r *http.Request) {
		server.MemoryDb(w,r)
	})

	port := GetEnvVariableOrDefault("PORT", DEFAULT_PORT)
	host := fmt.Sprintf("%s:%s", "", port)
	log.Printf("server setup complete")
	log.Printf("running on : %s", host)
	log.Fatal(http.ListenAndServe(host, nil))
}