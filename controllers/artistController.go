package controllers

import (
	"groupie_tracker/database"
	"groupie_tracker/models"
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
	if len(search) > 500 {
		ErrorController(w, r, http.StatusRequestEntityTooLarge)
		return
	}
	members := r.URL.Query()["member"]
	location := r.URL.Query().Get("location")
	if len(location) > 500 {
		ErrorController(w, r, http.StatusRequestEntityTooLarge)
		return
	}
	creationDateFrom := r.URL.Query().Get("creation_date_from")
	creationDateTo := r.URL.Query().Get("creation_date_to")
	albumDateFrom := r.URL.Query().Get("album_date_from")
	albumDateTo := r.URL.Query().Get("album_date_to")
	data := struct {
		Artists     []models.Artist
		Suggestions []models.Artist
	}{
		Suggestions: artists,
	}
	if len(members) > 0 || len(location) > 0 || len(creationDateFrom) > 0 || len(creationDateTo) > 0 || len(albumDateFrom) > 0 || len(albumDateTo) > 0 {
		filteredArtists, _ := Filter(artists, members, location, creationDateFrom, creationDateTo, albumDateFrom, albumDateTo)
		data.Artists = filteredArtists
		if len(data.Artists) == 0 {
			ErrorController(w, r, http.StatusNotFound)
			return
		}
		ParseController(w, r, "index", data)
	} else if len(search) > 0 {
		searchedArtists, _ := Search(artists, search)
		data.Artists = searchedArtists
		ParseController(w, r, "index", data)
	} else {
		data.Artists = artists
		ParseController(w, r, "index", data)
	}
}
