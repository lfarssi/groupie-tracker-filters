package controllers

import (
	"groupie_tracker/database"
	"html/template"
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
	members := r.URL.Query().Get("members")
    location := r.URL.Query().Get("location")
    creationDateFrom := r.URL.Query().Get("creation_date_from")
    creationDateTo := r.URL.Query().Get("creation_date_to")
    albumDateFrom := r.URL.Query().Get("album_date_from")
    albumDateTo := r.URL.Query().Get("album_date_to")
	if len(members) > 0  || len(location) > 0 || len(creationDateFrom) > 0 || len(creationDateTo) > 0 || len(albumDateFrom) > 0 || len(albumDateTo) > 0 {
		data, _ := filter(artists, members, location, creationDateFrom, creationDateTo, albumDateFrom, albumDateTo)

		res, err1 := template.ParseFiles("views/index.html")
		if err1 != nil {
			ErrorController(w, r, http.StatusInternalServerError)
			return
		}

		res.Execute(w, data)
	}
	if len(search) > 0 {
		data, _ := Search(artists, search)

		res, err1 := template.ParseFiles("views/index.html")
		if err1 != nil {
			ErrorController(w, r, http.StatusInternalServerError)
			return
		}

		res.Execute(w, data)
	} else {
		res, err1 := template.ParseFiles("views/index.html")
		if err1 != nil {
			ErrorController(w, r, http.StatusInternalServerError)
			return
		}
		
		if err := res.Execute(w, artists); err != nil {
			ErrorController(w, r, http.StatusInternalServerError)
			return
		}
	}


}
