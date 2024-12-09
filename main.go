package main

import (
	"fmt"
	"groupie_tracker/database"
	"groupie_tracker/models"
	"groupie_tracker/routes"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 1 {
		os.Exit(1)
	}
	var err error
	go func () {
		database.Artists, err = models.GetArtists()
	}()
	go func() {

		database.Locations, err = models.GetLocation()
	}()
	go func() {

		database.Dates, err = models.GetDate()
	}()
	for i := 0; i < len(database.Artists); i++ {
		artist := &database.Artists[i]
		if len(artist.Members) == 1 {
			artist.Type = "Artist"
		} else {
			artist.Type = "Group of " + strconv.Itoa(len(artist.Members))
		}
	}
	if err != nil {
		fmt.Println(err)
	}
	routes.Router()
}
