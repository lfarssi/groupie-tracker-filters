package controllers

import (
	"fmt"
	"groupie_tracker/models"
	"strconv"
	"strings"
)

func Filter(artists []models.Artist, members []string, location string, creationDateFrom string, creationDateTo string, albumDateFrom string, albumDateTo string) (any, error) {
	var filtered []models.Artist

	for _, artist := range artists {
		memberExists := len(members) == 0 // Default to true if no member filter
		locationExists := len(location) == 0
		creationDateFromExists := len(creationDateFrom) == 0
		creationDateToExists := len(creationDateTo) == 0
		albumDateFromExists := len(albumDateFrom) == 0
		albumDateToExists := len(albumDateTo) == 0
		albumYearStr := artist.FirstAlbum[len(artist.FirstAlbum)-4:] // Extract last 4 characters as year
		albumYear, err := strconv.Atoi(albumYearStr)
		if err != nil {
			return nil, fmt.Errorf("invalid FirstAlbum year format for artist %s: %v", artist.Name, err)
		}
		if len(members) > 0 {
			for _, member := range members {
				memberI, err := strconv.Atoi(string(member))
				if err != nil {
					return nil, err
				}
				if len(artist.Members) == memberI {
					memberExists = true

				}
			}
		} else {
			memberExists = true
		}
		if len(location) > 0 {
			for _, locations := range artist.Locationsr {
				if strings.Contains(locations, location) {
					locationExists = true

				}
			}
		}

		if len(creationDateFrom) > 0 && len(creationDateTo) > 0 {
			from, err := strconv.Atoi(creationDateFrom)
			if err != nil {
				return nil, err
			}
			to, err := strconv.Atoi(creationDateTo)
			if err != nil {
				return nil, err
			}
			if from < to {
				if artist.CreationDate > from  && artist.CreationDate < to {
					creationDateFromExists = true
				}
			} else {
				if artist.CreationDate > to  && artist.CreationDate < from {
                    creationDateFromExists = true
                }
			}

		}


		if len(albumDateFrom) > 0  && len(albumDateTo) > 0 {
			from, err := strconv.Atoi(albumDateFrom)
			if err != nil {
				return nil, fmt.Errorf("invalid albumDateFrom format: %v", err)
			}
			to, err := strconv.Atoi(albumDateTo)
			if err != nil {
				return nil, fmt.Errorf("invalid albumDateFrom format: %v", err)
			}
			if from < to {
				if albumYear > from  && albumYear < to {
					albumDateFromExists = true
				}
			} else {
				if albumYear > to  && albumYear < from {
					albumDateFromExists = true
				}
			}
			

		}


		if memberExists && locationExists && creationDateFromExists && creationDateToExists && albumDateFromExists && albumDateToExists {
			filtered = append(filtered, artist)
		}
	}
	return filtered, nil
}
