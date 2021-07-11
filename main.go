package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Camera struct {
	Name        string `json:"Name"`
	URL         string `json:"URL"`
	Description string `json:"Description"`
}

type Cameras []Camera

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<a href='/cameras'>Cameras</a>")
}

func allCameras(w http.ResponseWriter, r *http.Request) {
	cameras :=
		Cameras{
			Camera{Name: "Morskie Oko", URL: "http://pogoda.topr.pl/download/current/mors.jpeg", Description: "Morskie Oko view"},
			Camera{Name: "Morskie Oko", URL: "http://pogoda.topr.pl/download/current/mors.jpeg", Description: "Morskie Oko view"},
			Camera{Name: "Morskie Oko", URL: "http://pogoda.topr.pl/download/current/mors.jpeg", Description: "Morskie Oko view"},
		}
	fmt.Println("All cameras")
	json.NewEncoder(w).Encode(cameras)
}

func handleRequests() {
	http.HandleFunc("/", home)
	http.HandleFunc("/cameras", allCameras)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
