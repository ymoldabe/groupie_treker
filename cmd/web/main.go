package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

type Artist struct {
	ID            int      `json:"id"`
	IMAGE         string   `json:"image"`
	NAME          string   `json:"name"`
	MEMBERS       []string `json:"members"`
	CREATION_DATE int      `json:"creationDate"`
	FIRST_ALBUM   string   `json:"firstAlbum"`
	LOCATIONS     string   `json:"locations"`
	CONCERT_DATES string   `json:"concertDates"`
	RELATIONS     string   `json:"relations"`
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

func artist(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		Error(w, http.StatusBadRequest)
		return
	}
	checkID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || checkID < 1 {
		Error(w, http.StatusNotFound)
		return
	}
	id := strconv.Itoa(checkID)
	artistData := Artist{}
	urlWay := "https://groupietrackers.herokuapp.com/api/artists" + "/" + id
	json4ik, err := http.Get(urlWay)
	if err != nil {
		Error(w, http.StatusInternalServerError)
		return
	}
	defer json4ik.Body.Close()

	body, err := ioutil.ReadAll(json4ik.Body)
	if err != nil {
		Error(w, http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal([]byte(body), &artistData)
	if err != nil {
		Error(w, http.StatusInternalServerError)
		return
	}
	files := []string{
		"/home/student/groupie_treker/ui/html/artistData.html",
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		Error(w, http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, artistData)
}

func group(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		Error(w, http.StatusBadRequest)
		return
	}
	if r.URL.Path != "/" {
		Error(w, 404)
		return
	}

	groups := []Artist{}

	urlWay := "https://groupietrackers.herokuapp.com/api/artists"

	json4ik, err := http.Get(urlWay)
	if err != nil {
		Error(w, http.StatusInternalServerError)
		return
	}
	defer json4ik.Body.Close()

	body, err := ioutil.ReadAll(json4ik.Body)
	if err != nil {
		Error(w, http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal([]byte(body), &groups)
	if err != nil {
		Error(w, http.StatusInternalServerError)
		return
	}
	files := []string{
		"/home/student/groupie_treker/ui/html/body_home.html",
		"/home/student/groupie_treker/ui/html/footer_partial.html",
		"/home/student/groupie_treker/ui/html/front.base.html",
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		Error(w, http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, groups)
}
