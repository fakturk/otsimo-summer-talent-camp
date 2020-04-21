package main

import(
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main()  {
	router:=mux.NewRouter()
	router.HandleFunc("/",helloFunc).Methods("GET")
	http.ListenAndServe(":8080",router)

}

func helloFunc(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer,"Hello Otsimo\n")
}
