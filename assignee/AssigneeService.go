package assignee

import (
	"context"
	"fmt"
	"github.com/fakturk/otsimo-summer-talent-camp/db"
	"github.com/fakturk/otsimo-summer-talent-camp/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindAssigneeIDByName(name string) (string, error) {
	var assignee model.Assignee

	collection := db.ConnectDB("Assignees")

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}

	//we make our search case insensitive by using regex
	filter := bson.M{"name": primitive.Regex{Pattern: name, Options: ""}}
	err := collection.FindOne(context.TODO(), filter).Decode(&assignee)
	fmt.Println(assignee)
	return assignee.ID,err
}

