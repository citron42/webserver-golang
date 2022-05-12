// Simple web server in golang
// Note trad : Handler = module

package main

import (
	"fmt"
	"log"      // use to print fatal error in case there is one
	"net/http" // use to provide all the functionality for creating HTTP client
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound) // Status not found corresponds to a 404 error
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello !")
}

func formHandler(w http.ResponseWriter, r *http.Request) { // only for form if you want to use it
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static")) // create a new file server which will serve files from the current directory

	// Don't forget to add http.Handle() line whenever you add a new webpage

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil { // first parameters of ListenAndServer method need to be in string
		log.Fatal(err)
	}
}
