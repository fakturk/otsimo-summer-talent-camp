package main

import (
	"fmt"
	"github.com/fakturk/otsimo-summer-talent-camp/assignee"
	"github.com/fakturk/otsimo-summer-talent-camp/candidate"
	"github.com/gorilla/mux"
	"net/http"
)



func main()  {
	router:=mux.NewRouter()
	router.HandleFunc("/",helloFunc).Methods("GET")
	router.HandleFunc("/candidate/create",candidate.CreateCandidateFunc).Methods("POST")
	router.HandleFunc("/candidate/read/{id}",candidate.ReadCandidateFunc).Methods("GET")
	router.HandleFunc("/candidate/delete/{id}",candidate.DeleteCandidateFunc).Methods("DELETE")
	router.HandleFunc("/candidate/accept/{id}",candidate.AcceptCandidateFunc).Methods("GET")
	router.HandleFunc("/candidate/deny/{id}",candidate.DenyCandidateFunc).Methods("GET")
	router.HandleFunc("/candidates",candidate.GetCandidatesFunc).Methods("GET")
	router.HandleFunc("/meeting/arrange/{id}",candidate.ArrangeMeetingFunc).Methods("POST")
	router.HandleFunc("/meeting/complete/{id}",candidate.CompleteMeetingFunc).Methods("GET")
	router.HandleFunc("/assignee/findid/{name}",assignee.FindAssigneeIDByNameFunc).Methods("GET")

	http.ListenAndServe(":8080",router)

}



func helloFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Hello Otsimo\n")
}
// ConnectDB : This is helper function to connect mongoDB
// If you want to export your function. You must to start upper case function name. Otherwise you won't see your function when you import that on other class.
