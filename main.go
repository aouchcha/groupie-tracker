package main

import (
	"fmt"
	"net/http"

	f "groupie-tracker/functions"
)

// Define a struct that matches the JSON structure

func main() {
	f.GetArtistData()
	f.GetRelationData()
	f.GetLocationData()
	f.GetDatesData()
	// http.Handle("/styles/",http.StripPrefix("/styles/", http.FileServer(http.Dir("./styles"))))
	http.HandleFunc("/styles/", f.ServeStyle)
	http.HandleFunc("/", f.FirstPage)
	http.HandleFunc("/artist", f.OtherPages)
	fmt.Println("http://localhost:8081")
	http.ListenAndServe(":8081", nil)
}
