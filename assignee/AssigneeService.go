package assignee

import (
	"context"
	"fmt"
	"github.com/fakturk/otsimo-summer-talent-camp/db"
	"github.com/fakturk/otsimo-summer-talent-camp/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func FindAssigneeIDByName(name string) (string, error) {
	var assignee model.Assignee

	collection := db.ConnectDB("Assignees")

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}

	//we make our search case insensitive by using regex
	filter := bson.M{"name": primitive.Regex{Pattern: name, Options: ""}}
	err := collection.FindOne(context.TODO(), filter).Decode(&assignee)
	//fmt.Println(assignee)
	return assignee.ID,err
}

func GetAssignee(_id string) (model.Assignee, error) {
	var assignee model.Assignee

	collection := db.ConnectDB("Assignees")
	filter := bson.M{"_id": _id}
	err := collection.FindOne(context.TODO(), filter).Decode(&assignee)

	return assignee,err
}

func FindAssigneesCandidates(_id string) ([]model.Candidate, error) {
	var candidates []model.Candidate
	//Connection mongoDB with helper class
	collection := db.ConnectDB("Candidates")
	fmt.Println(collection)
	// bson.M{},  we passed empty filter. So we want to get all data.
	cur, err := collection.Find(context.TODO(), bson.M{})


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
		//fmt.Println(candidate)
		// add item our array
		if candidate.Assignee==_id {
			candidates = append(candidates, candidate)
		}

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return candidates,err

}
