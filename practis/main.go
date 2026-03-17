package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Printf("hello")

	fileserver := http.FileServer(http.Dir("../static"))
	http.Handle("/", fileserver)

	fmt.Printf(" \n Starting the server <3  ( ctrl + c to exit ) \n")
	if err := http.ListenAndServe(":5040", nil); err != nil {

		log.Fatal(err)

	}

}
