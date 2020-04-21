package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Candidate struct {
	ID string  `json:"_id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Department string `json:"department"`
	University string `json:"university"`
	Experience bool `json:"experience"`
	Status string `json:"status"`
	MeetingCount int `json:"meeting_count"`
	NextMeeting string `json:"next_meeting"`
	Assignee string `json:"assignee"`
}

func main()  {
	router:=mux.NewRouter()
	router.HandleFunc("/",helloFunc).Methods("GET")
	router.HandleFunc("candidate/create",createCandidate).Methods("POST")
	http.ListenAndServe(":8080",router)

}

func createCandidate(w http.ResponseWriter, r *http.Request) {

}

func helloFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Hello Otsimo\n")
}
