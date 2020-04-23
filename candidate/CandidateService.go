package candidate

import (
	"context"
	"github.com/fakturk/otsimo-summer-talent-camp/db"
	"github.com/fakturk/otsimo-summer-talent-camp/model"
	"go.mongodb.org/mongo-driver/mongo"
)


func CreateCandidate(candidate model.Candidate) (model.Candidate, *mongo.InsertOneResult, error) {
	// connect db
	collection := db.ConnectDB()

	// insert our book model.
	result, err := collection.InsertOne(context.TODO(), candidate)

	return candidate,result,err
}
