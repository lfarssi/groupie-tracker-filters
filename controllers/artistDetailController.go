package controllers

import (
	"encoding/json"
	"groupie_tracker/models"
	"html/template"
	"net/http"
	"strconv"
	"sync"
)

func fetchData(url string, target interface{}) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if err = json.NewDecoder(response.Body).Decode(target); err != nil {
		return err
	}
	return nil
}

func ArtistDetailController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorController(w, r, http.StatusMethodNotAllowed)
		return
	}
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ErrorController(w, r, http.StatusBadRequest)
		return
	}
	artist, err := models.ArtistById(id)
	if err != nil {
		ErrorController(w, r, http.StatusNotFound)
		return
	}

	var locations models.Location
	var dates models.Date
	var relations models.Relation

	channels := make(chan error, 3)
	var wg sync.WaitGroup
	wg.Add(3) 

	go func() {
		defer wg.Done()
		channels <- fetchData(artist.Locations, &locations)
	}()
	go func() {
		defer wg.Done()
		channels <- fetchData(artist.ConcertDates, &dates)
	}()
	go func() {
		defer wg.Done()
		channels <- fetchData(artist.Relations, &relations)
	}()

	go func() {
		wg.Wait()      
		close(channels) 
	}()

	for err := range channels {
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError)
			return
		}
	}

	res, err := template.ParseFiles("views/detail.html")
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError)
		return
	}
	artistDetail := models.Artist{
		Id:           artist.Id,
		Name:         artist.Name,
		Image:        artist.Image,
		Members:      artist.Members,
		CreationDate: artist.CreationDate,
		FirstAlbum:   artist.FirstAlbum,
		Location:    locations.Location,
		ConcertDate: dates.Dates,
		Relation:     relations.DatesLocations,
	}
	res.Execute(w, artistDetail)
}
