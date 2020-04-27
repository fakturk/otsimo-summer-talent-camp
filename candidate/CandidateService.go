package candidate

import (
	"context"
	"fmt"
	"github.com/fakturk/otsimo-summer-talent-camp/db"
	"github.com/fakturk/otsimo-summer-talent-camp/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

func GetAllCandidates() ([]model.Candidate, error) {
	var candidates []model.Candidate
	//Connection mongoDB with helper class
	collection := db.ConnectDB()
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
		fmt.Println(candidate)
		// add item our array
		candidates = append(candidates, candidate)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return candidates,err
}
func CreateCandidate(candidate model.Candidate) (model.Candidate, *mongo.InsertOneResult, error) {
	// connect db
	collection := db.ConnectDB()

	//create candidate unique id
	candidate.ID=primitive.NewObjectID().Hex()

	// insert our book model.
	result, err := collection.InsertOne(context.TODO(), candidate)

	return candidate,result,err
}

func ReadCandidate(_id string) (model.Candidate, error){
	var candidate model.Candidate

	collection := db.ConnectDB()

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": _id}
	err := collection.FindOne(context.TODO(), filter).Decode(&candidate)
	return candidate,err
}

func DeleteCandidate(_id string) (*mongo.DeleteResult, error) {

	collection := db.ConnectDB()

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": _id}
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	fmt.Println(deleteResult)
	return deleteResult,err
}

func ArrangeMeeting(_id string, meetingTime *time.Time) (*mongo.UpdateResult, error) {
	collection := db.ConnectDB()
	var candidate model.Candidate

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": _id}
	err := collection.FindOne(context.TODO(), filter).Decode(&candidate)
	candidate.Next_Meeting=meetingTime
	updateResult,err:=collection.ReplaceOne(context.TODO(),filter,candidate)

	return updateResult,err
}
//func CompleteMeeting (_id string) error{
//
//}
//
func DenyCandidate(_id string) (*mongo.UpdateResult, error) {
	collection := db.ConnectDB()
	var candidate model.Candidate

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": _id}
	err := collection.FindOne(context.TODO(), filter).Decode(&candidate)
	candidate.Status="Denied"
	updateResult,err:=collection.ReplaceOne(context.TODO(),filter,candidate)

	return updateResult,err
}

func AcceptCandidate(_id string) (*mongo.UpdateResult, error) {
	collection := db.ConnectDB()
	var candidate model.Candidate

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": _id}
	err := collection.FindOne(context.TODO(), filter).Decode(&candidate)
	candidate.Status="Accepted"
	updateResult,err:=collection.ReplaceOne(context.TODO(),filter,candidate)

	return updateResult,err
}
//func FindAssigneeIDByName (name string) string{
//
//}
