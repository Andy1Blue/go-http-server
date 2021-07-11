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
	http.HandleFunc("/foobars", allFoobars)
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), nil))
}

func main() {
	handleRequests()
}
