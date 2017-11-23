package main

import (
	"fmt"
	"net/http"

	"./src/util"
)

//can name it anything you want i.e handlerFunc
func handlerFunc(w http.ResponseWriter, r *http.Request) {

	// //get the string from the input box
	//userSent := r.Header.Get("userAskEliza")
	 userSent := r.URL.Query().Get("value")

	//send the answer to the user
	fmt.Fprintf(w, "\nEliza: %s\n", util.ReplyQuestion(userSent))

}

func main() {

	//opened  xhr.open("POST", "/askEliza"); and set the handler here
	http.HandleFunc("/user-input", handlerFunc)

	//serves the folder
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8080", nil)

}
