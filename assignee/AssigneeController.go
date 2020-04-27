package assignee

import (
	"encoding/json"
	"github.com/fakturk/otsimo-summer-talent-camp/helper"
	"github.com/gorilla/mux"
	"net/http"
)
func FindAssigneeIDByNameFunc(w http.ResponseWriter, r *http.Request) {
	// we get params with mux.
	var params = mux.Vars(r)
	name := params["name"]
	result,err :=FindAssigneeIDByName(name)
	if err != nil {
		helper.GetError(err, w)
		return

	}

	json.NewEncoder(w).Encode(result)
}

func FindAssigneesCandidatesFunc(w http.ResponseWriter, r *http.Request) {
	// we get params with mux.
	var params = mux.Vars(r)
	id := params["id"]
	result,err :=FindAssigneesCandidates(id)
	if err != nil {
		helper.GetError(err, w)
		return

	}

	json.NewEncoder(w).Encode(result)
}

