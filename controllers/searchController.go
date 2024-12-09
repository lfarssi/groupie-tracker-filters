package controllers

import (
	"groupie_tracker/models"
	"strconv"
	"strings"
)

func Search(artists []models.Artist, search string) ([]models.Artist, error) {
	search = strings.ToLower(strings.TrimSpace(search))
	if search == "" {
		// Return all artists if search is empty, or handle differently as per requirements
		return artists, nil
	}

	searched := make([]models.Artist, 0)
	uniqueArtists := make(map[int]struct{}) // Track unique artists by their ID

	for _, artist := range artists {
		// Skip if artist is already added
		if _, exists := uniqueArtists[artist.Id]; exists {
			continue
		}

		// Match criteria
		isMatch := false

		// Check artist name, first album, or creation date
		if strings.Contains(strings.ToLower(artist.Name), search) ||
			strings.Contains(strings.ToLower(artist.FirstAlbum), search) ||
			strings.Contains(strconv.Itoa(artist.CreationDate), search) {
			isMatch = true
		}

		// Check members
		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), search) {
				isMatch = true
				break
			}
		}

		// Check concert dates
		for _, date := range artist.ConcertDatesr {
			if strings.Contains(strings.ToLower(date), search) {
				isMatch = true
				break
			}
		}

		// Check locations
		for _, location := range artist.Locationsr {
			if strings.Contains(strings.ToLower(location), search) {
				isMatch = true
				break
			}
		}

		// Add to results if a match is found
		if isMatch {
			searched = append(searched, artist)
			uniqueArtists[artist.Id] = struct{}{} // Mark this artist as added
		}
	}

	return searched, nil
}
