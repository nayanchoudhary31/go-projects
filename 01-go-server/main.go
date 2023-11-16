package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parse Form Error: %v\n", err)
		return
	}
	fmt.Fprintf(w, "Post Request Successfully\n")
	name := r.FormValue("name")
	fmt.Print(name)
	email := r.FormValue("email")

	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Email: %s\n", email)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")

}
