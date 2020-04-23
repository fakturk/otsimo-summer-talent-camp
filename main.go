package main

import (
	"fmt"
	"github.com/fakturk/otsimo-summer-talent-camp/candidate"
	"github.com/fakturk/otsimo-summer-talent-camp/db"
	"github.com/gorilla/mux"
	"net/http"
)



func main()  {
	router:=mux.NewRouter()
	db.ConnectDB()
	router.HandleFunc("/",helloFunc).Methods("GET")
	router.HandleFunc("/candidate/create",candidate.CreateCandidateFunc).Methods("POST")
	router.HandleFunc("/candidates",candidate.GetCandidatesFunc).Methods("GET")
	http.ListenAndServe(":8080",router)

}



func helloFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Hello Otsimo\n")
}
// ConnectDB : This is helper function to connect mongoDB
// If you want to export your function. You must to start upper case function name. Otherwise you won't see your function when you import that on other class.
