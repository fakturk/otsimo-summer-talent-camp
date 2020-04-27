package candidate

import (
	"context"
	"errors"
	"fmt"
	"github.com/fakturk/otsimo-summer-talent-camp/assignee"
	"github.com/fakturk/otsimo-summer-talent-camp/db"
	"github.com/fakturk/otsimo-summer-talent-camp/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"regexp"
	"time"
)

func GetAllCandidates() ([]model.Candidate, error) {
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
	collection := db.ConnectDB("Candidates")
	var err error
	var result *mongo.InsertOneResult
	var candidatesAssignee model.Assignee

	//if the candidates department is not marketing, design or development, we return an error
	if !(candidate.Department=="Marketing" ||  candidate.Department=="Design" || candidate.Department=="Development") {
		//fmt.Println("inside department error")
		err= errors.New("Deparment should be Marketing, Design or Development")
	}

	//check assignee department
	candidatesAssignee,err=assignee.GetAssignee(candidate.Assignee)
	if candidate.Department!=candidatesAssignee.Department{
		err= errors.New("New candidates should have an assignee who is working in the department that the candidate is applying to work.")
	}

	//validate email according to HTML specification https://html.spec.whatwg.org/multipage/input.html#valid-e-mail-address
	//However this validation specification did not forcifully use a@b.c format instead it uses a@b format , dot is not a must in their document
	//because of it I changed the regex that accept only a@b.c format
	var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])*$")
	if !rxEmail.MatchString(candidate.Email) {
		err=errors.New("Not a valid email address")

	}
	if candidate.Status!="Pending"{
		candidate.Status="Pending"
	}
	if candidate.Meeting_Count!=0{
		candidate.Meeting_Count=0
	}


	//create candidate unique id
	candidate.ID=primitive.NewObjectID().Hex()

	// insert our book model.
	if err==nil {
		result, err = collection.InsertOne(context.TODO(), candidate)
	}

	return candidate,result,err
}

func ReadCandidate(_id string) (model.Candidate, error){
	var candidate model.Candidate

	collection := db.ConnectDB("Candidates")

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": _id}
	err := collection.FindOne(context.TODO(), filter).Decode(&candidate)
	return candidate,err
}

func DeleteCandidate(_id string) (*mongo.DeleteResult, error) {

	collection := db.ConnectDB("Candidates")

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": _id}
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	fmt.Println(deleteResult)
	return deleteResult,err
}

func ArrangeMeeting(_id string, meetingTime *time.Time) (*mongo.UpdateResult, error) {
	collection := db.ConnectDB("Candidates")
	var candidate model.Candidate

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": _id}
	err := collection.FindOne(context.TODO(), filter).Decode(&candidate)
	candidate.Next_Meeting=meetingTime
	//if candidates next meeting is their 4th meeting (or in another words if their meeting count is 3)
	//their assigne changed to Zafer who is the CEO
	if candidate.Meeting_Count==3 {
		candidate.Assignee="Zafer"
	}
	updateResult,err:=collection.ReplaceOne(context.TODO(),filter,candidate)

	return updateResult,err
}
//increase meeting count by 1 and make next meeting null
func CompleteMeeting(_id string) (*mongo.UpdateResult, error) {
	collection := db.ConnectDB("Candidates")
	var candidate model.Candidate

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": _id}
	err := collection.FindOne(context.TODO(), filter).Decode(&candidate)
	candidate.Meeting_Count+=1
	if candidate.Meeting_Count>0&&candidate.Meeting_Count<4 {
		candidate.Status="In Progress"
	}
	candidate.Next_Meeting=nil
	updateResult,err:=collection.ReplaceOne(context.TODO(),filter,candidate)

	return updateResult,err
}

func DenyCandidate(_id string) (*mongo.UpdateResult, error) {
	collection := db.ConnectDB("Candidates")
	var candidate model.Candidate

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": _id}
	err := collection.FindOne(context.TODO(), filter).Decode(&candidate)
	candidate.Status="Denied"
	updateResult,err:=collection.ReplaceOne(context.TODO(),filter,candidate)

	return updateResult,err
}

func AcceptCandidate(_id string) (*mongo.UpdateResult, error) {
	collection := db.ConnectDB("Candidates")
	var candidate model.Candidate

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": _id}
	var err error
	var updateResult *mongo.UpdateResult
	err = collection.FindOne(context.TODO(), filter).Decode(&candidate)
	if candidate.Meeting_Count<4 {
		err = errors.New("Candidates cannot be accepted before the completion of 4 meetings.")
	}
	if err==nil {
		candidate.Status="Accepted"
		updateResult,err=collection.ReplaceOne(context.TODO(),filter,candidate)
	}


	return updateResult,err
}
//func FindAssigneeIDByName (name string) string{
//
//}
