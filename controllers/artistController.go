package controllers

import (
	"fmt"
	"groupie_tracker/database"
	"net/http"
)

func ArtistController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorController(w, r, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		ErrorController(w, r, http.StatusNotFound)
		return
	}
	artists := database.Artists
	locations := database.Locations
	dates := database.Dates
	for i, _ := range artists {
		artists[i].Locationsr = locations.Index[i].Location
		artists[i].ConcertDatesr = dates.Index[i].Dates
	}
	search := r.URL.Query().Get("search")
	members := r.URL.Query().Get("member")
	location := r.URL.Query().Get("location")
	creationDateFrom := r.URL.Query().Get("creation_date_from")
	creationDateTo := r.URL.Query().Get("creation_date_to")
	albumDateFrom := r.URL.Query().Get("album_date_from")
	albumDateTo := r.URL.Query().Get("album_date_to")
	fmt.Println(members)
	if len(members) > 0 || len(location) > 0 || len(creationDateFrom) > 0 || len(creationDateTo) > 0 || len(albumDateFrom) > 0 || len(albumDateTo) > 0 {
		data, _ := Filter(artists, members, location, creationDateFrom, creationDateTo, albumDateFrom, albumDateTo)
		ParseController(w, r, "index", data)
	} else if len(search) > 0 {
		data, _ := Search(artists, search)
		ParseController(w, r, "index", data)
	} else {
		ParseController(w, r, "index", artists)
	}
}
