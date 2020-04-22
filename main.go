package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

type Candidate struct {
	ID string  `json:"_id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Department string `json:"department"`
	University string `json:"university"`
	Experience bool `json:"experience"`
	Status string `json:"status"`
	MeetingCount int `json:"meeting_count"`
	NextMeeting string `json:"next_meeting"`
	Assignee string `json:"assignee"`
}

type Assignee struct {
	ID string  `json:"_id"`
	Name string `json:"name"`
	Department string `json:"department"`
}

func main()  {
	router:=mux.NewRouter()
	ConnectDB()
	router.HandleFunc("/",helloFunc).Methods("GET")
	router.HandleFunc("/candidate/create",createCandidate).Methods("POST")
	router.HandleFunc("/candidates",getCandidates).Methods("GET")
	http.ListenAndServe(":8080",router)

}

func getCandidates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var candidates []Candidate
	//Connection mongoDB with helper class
	collection := ConnectDB()
	fmt.Println(collection)
	// bson.M{},  we passed empty filter. So we want to get all data.
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		fmt.Println(err)
		return
	}

	// Close the cursor once finished
	/*A defer statement defers the execution of a function until the surrounding function returns.
	simply, run cur.Close() process but after cur.Next() finished.*/
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var candidate Candidate
		// & character returns the memory address of the following variable.
		err := cur.Decode(&candidate) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		candidates = append(candidates, candidate)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(candidates) // encode similar to serialize process.
}

func createCandidate(w http.ResponseWriter, r *http.Request) {

}

func helloFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Hello Otsimo\n")
}
// ConnectDB : This is helper function to connect mongoDB
// If you want to export your function. You must to start upper case function name. Otherwise you won't see your function when you import that on other class.
func ConnectDB() *mongo.Collection {

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("Otsimo").Collection("Candidates")

	return collection
}