package candidate

import (
	"context"
	"encoding/json"
	"github.com/fakturk/otsimo-summer-talent-camp/db"
	"github.com/fakturk/otsimo-summer-talent-camp/helper"
	"github.com/fakturk/otsimo-summer-talent-camp/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func GetCandidatesFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	candidates,err:= GetAllCandidates()
	if err != nil {
		helper.GetError(err, w)

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

func ReadCandidateFunc(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")

	// we get params with mux.

	var params = mux.Vars(r)
	id := params["id"]

	candidate,err:=ReadCandidate(id)
	if err != nil {
		helper.GetError(err, w)
		return

	}

	json.NewEncoder(w).Encode(candidate)

}

func DeleteCandidateFunc(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")

	// we get params with mux.
	var params = mux.Vars(r)
	id := params["id"]

	result, err :=DeleteCandidate(id)
	if err != nil {
		helper.GetError(err, w)
		return

	}

	json.NewEncoder(w).Encode(result)

}

func AcceptCandidateFunc(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")

	// we get params with mux.
	var params = mux.Vars(r)
	id := params["id"]

	result,err :=AcceptCandidate(id)
	if err != nil {
		helper.GetError(err, w)
		return

	}

	json.NewEncoder(w).Encode(result)

}



func DenyCandidateFunc(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")

	// we get params with mux.
	var params = mux.Vars(r)
	id := params["id"]

	result,err :=DenyCandidate(id)
	if err != nil {
		helper.GetError(err, w)
		return

	}

	json.NewEncoder(w).Encode(result)

}

