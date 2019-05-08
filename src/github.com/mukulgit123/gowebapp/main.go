package main


import (
	"fmt"
	
	"strings"

	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r :=  mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	r.HandleFunc("/bye", handler).Methods("GET")
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// For this case, we will always pipe "Hello World" into the response writer
	//fmt.Fprintf(w,r)
	if strings.Contains(string(r.URL.Path),"bye")	{
	fmt.Fprintf(w, "Good Bye!\n")
	} else	{
	fmt.Fprintf(w,"Hello World!\n")
	}
}
