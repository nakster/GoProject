package main

import (
	"net/http"
	"fmt"

)

//can name it anything you want i.e handlerFunc
func handlerFunc(w http.ResponseWriter, r *http.Request) {

	// //get the string from the input box
	 userSent := r.Header.Get("userAskEliza")
	
	//send the answer to the user
	fmt.Fprintf(w, "\nEliza: %s\n", userSent)
	
}

func main() {

	//opened  xhr.open("POST", "/askEliza"); and set the handler here
	http.HandleFunc("/askEliza", handlerFunc)

	//serves the folder
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8080", nil)
	
}