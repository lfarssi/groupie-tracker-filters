package controllers

import (
	"groupie_tracker/models"
	"strconv"
	"strings"
)

func Search(artists []models.Artist, search string) ([]models.Artist, error) {
	search = strings.ToLower(search)
	searched := make([]models.Artist, 0)
	uniqueArtists := make(map[int]struct{}) // To track unique artists by their Id
	var kayn bool

	for _, artist := range artists {
		// Skip if artist has already been added
		if _, exists := uniqueArtists[artist.Id]; exists {
			continue
		}

		// Check if artist matches search criteria
		if strings.Contains(strings.ToLower(artist.Name), search) ||
			strings.Contains(artist.FirstAlbum, search) ||
			strings.Contains(strconv.Itoa(artist.CreationDate), search) {
			kayn = true
			uniqueArtists[artist.Id] = struct{}{}
			continue
		}

		// Check if any member matches the search criteria
		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), search) {
				kayn = true
				uniqueArtists[artist.Id] = struct{}{}
				break
			}
		}

		// Check if any concert date matches the search criteria
		for _, date := range artist.ConcertDatesr {
			if strings.Contains(strings.ToLower(date), search) {
				kayn = true
				uniqueArtists[artist.Id] = struct{}{}
				break
			}
		}

		// Check if any location matches the search criteria
		for _, location := range artist.Locationsr {
			if strings.Contains(strings.ToLower(location), search) {
				kayn = true
				uniqueArtists[artist.Id] = struct{}{}
				break
			}
		}
		if kayn {
			searched = append(searched, artist)
		}
		kayn = false
	}

	return searched, nil
}
