package main

import (
	"fmt"
	"log"
	"net/http"
)

func HellowHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/hello" {

		http.Error(w, " server not found ", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, " Methode is not found ", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "!hello")
}

// func HellowHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/hello" {
// 		http.Error(w, "4040 not found ", http.StatusNotFound)
// 		return
// 	}

// 	if r.Method != "GET" {
// 		http.Error(w, "method is not supported", http.StatusNotFound)
// 		return
// 	}
// 	fmt.Fprintf(w, "hello!")

// }

func main() {

	fmt.Printf("hello")

	fileserver := http.FileServer(http.Dir("../static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/hello", HellowHandler)

	fmt.Printf(" \n Starting the server <3  ( ctrl + c to exit ) \n")
	if err := http.ListenAndServe(":4040", nil); err != nil {

		log.Fatal(err)

	}

}
