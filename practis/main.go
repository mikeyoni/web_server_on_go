package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/charmbracelet/lipgloss"
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

// colure gui render in the termianls

// 1. Define your "Style Functions" globally
var (
	errorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Bold(true)
	userStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#00D7FF")).Italic(true)
	lineStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#555555"))
)

var (
	cynestyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#0084ff")).Bold(true).Align(lipgloss.Center)
)

// 2. Create a helper function to wrap the logic
func printUserMessage(name string, msg string) {
	// Every time this runs, it applies the userStyle to the name
	fmt.Printf("[%s] %s\n", userStyle.Render(name), msg)
}

func printDivider() {
	// This function changes the style for just one line
	fmt.Println(lineStyle.Render("----------------------------------"))
}

func main() {

	fmt.Println(cynestyle.Render("Hello, Pirate King! "))

	fmt.Printf("hello")

	// exampolsss

	// Now your main logic is clean and easy to read
	printUserMessage("Mikey", "The server is starting...")

	printDivider() // Next line changes style automatically

	fmt.Println(errorStyle.Render("ALERT: Connection from Phone detected!"))

	printDivider()

	// exampols

	fileserver := http.FileServer(http.Dir("../static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/hello", HellowHandler)

	fmt.Printf(" \n Starting the server <3  ( ctrl + c to exit ) \n")
	if err := http.ListenAndServe(":4040", nil); err != nil {

		log.Fatal(err)

	}

}
