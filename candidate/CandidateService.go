package candidate

import (
	"context"
	"github.com/fakturk/otsimo-summer-talent-camp/db"
	"github.com/fakturk/otsimo-summer-talent-camp/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


func CreateCandidate(candidate model.Candidate) (model.Candidate, *mongo.InsertOneResult, error) {
	// connect db
	collection := db.ConnectDB()

	//create candidate unique id
	candidate.ID=primitive.NewObjectID().Hex()
	
	// insert our book model.
	result, err := collection.InsertOne(context.TODO(), candidate)

	return candidate,result,err
}

//func ReadCandidate(_id string) (model.Candidate, error){
//
//}
//
//func DeleteCandidate(_id string) error{
//
//}
//
//func ArrangeMeeting(_id string, nextMeetingTime *time.Time) error{
//
//}
//func CompleteMeeting (_id string) error{
//
//}
//
//func DenyCandidate (_id string) error{
//
//}
//
//func AcceptCandidate(_id string) error{
//
//}
//
//func FindAssigneeIDByName (name string) string{
//
//}
