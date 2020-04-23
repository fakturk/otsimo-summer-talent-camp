package candidate

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fakturk/otsimo-summer-talent-camp/db"
	"github.com/fakturk/otsimo-summer-talent-camp/helper"
	"github.com/fakturk/otsimo-summer-talent-camp/model"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

func GetCandidatesFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var candidates []model.Candidate
	//Connection mongoDB with helper class
	collection := db.ConnectDB()
	fmt.Println(collection)
	// bson.M{},  we passed empty filter. So we want to get all data.
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	// Close the cursor once finished
	/*A defer statement defers the execution of a function until the surrounding function returns.
	simply, run cur.Close() process but after cur.Next() finished.*/
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var candidate model.Candidate
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

func CreateCandidateFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var candidate model.Candidate

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&candidate)

	_, result, err :=CreateCandidate(candidate)
	if err != nil {
		helper.GetError(err, w)
		return
	}
	json.NewEncoder(w).Encode(result)
}
