package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Foobar struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

type Foobars []Foobar

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello GO")
}

func foobar(w http.ResponseWriter, r *http.Request) {
	foobar := Foobar{Name: "Random 1", Description: "Random 2"}
	fmt.Println("Foobar")
	json.NewEncoder(w).Encode(foobar)
}

func allFoobars(w http.ResponseWriter, r *http.Request) {
	foobars :=
		Foobars{
			Foobar{Name: "Random 1", Description: "Random 2"},
			Foobar{Name: "Random 2", Description: "Random 4"},
		}
	fmt.Println("All foobars")
	json.NewEncoder(w).Encode(foobars)
}

func handleRequests() {
	http.HandleFunc("/", home)
	http.HandleFunc("/foobar", foobar)
	http.HandleFunc("/foobars", allFoobars)

	port := os.Getenv("PORT")

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func main() {
	handleRequests()
}
