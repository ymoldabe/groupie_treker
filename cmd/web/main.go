package main

import (
	"log"
	"net/http"
)

type Artist struct {
	ID                 int                 `json:"id"`
	IMAGE              string              `json:"image"`
	NAME               string              `json:"name"`
	MEMBERS            []string            `json:"members"`
	CREATION_DATE      int                 `json:"creationDate"`
	FIRST_ALBUM        string              `json:"firstAlbum"`
	LOCATIONS          string              `json:"locations"`
	CONCERT_DATES      string              `json:"concertDates"`
	RELATIONS          string              `json:"relations"`
	LOCATION_AND_DATES map[string][]string `json:"datesLocations"`
}

func main() {
	mux := http.NewServeMux()
	styles := http.FileServer(http.Dir("/home/student/groupie_treker/ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", styles))
	mux.HandleFunc("/", group)
	mux.HandleFunc("/artist", artist)
	log.Println("Go to run http://localhost:8000/")
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}
