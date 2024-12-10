package controllers

import (
	"fmt"
	"groupie_tracker/models"
	"strconv"
	"strings"
)

// func isInRange(value, from, to int) bool {
// 	if from > to {
// 		from, to = to, from
// 		return value >= to && value <= to
// 	}
// 	return value >= from && value <= to
// }

func Filter(artists []models.Artist, members []string, location string, creationDateFrom string, creationDateTo string, albumDateFrom string, albumDateTo string) ([]models.Artist, error) {
	var filtered []models.Artist

	for _, artist := range artists {
		memberExists := len(members) == 0 // Default to true if no member filter
		locationExists := len(location) == 0
		creationDateInRange := len(creationDateFrom) == 0
		albumDateInRange := len(albumDateFrom) == 0
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
					break
				} else {
					memberExists = false
				}
			}
		} else {
			memberExists = true
		}
		if len(location) > 0 {
			for _, locations := range artist.Locationsr {
				if strings.Contains(locations, location) {
					locationExists = true
					break
				} else {
					locationExists = false
				}
			}
		}
		if len(creationDateFrom) > 0 && len(creationDateTo) > 0 {
			from, err := strconv.Atoi(creationDateFrom)
			if err != nil {
				return nil, fmt.Errorf("invalid creationDateFrom format: %v", err)
			}
			to, err := strconv.Atoi(creationDateTo)
			if err != nil  {
				return nil, fmt.Errorf("invalid creationDateTo format: %v", err)
			}
			if (from < 1950 || from > 2025) || (to < 1950 || to > 2025) {
                return nil, fmt.Errorf("invalid creationDateTo format: %v", err)
            }
			
			if artist.CreationDate >= from && artist.CreationDate <= to {
				creationDateInRange = true
			} else {
				creationDateInRange = false
			}
		}

		if len(albumDateFrom) > 0 && len(albumDateTo) > 0 {
			from, err := strconv.Atoi(albumDateFrom)
			if err != nil {
				return nil, fmt.Errorf("invalid albumDateFrom format: %v", err)
			}
			to, err := strconv.Atoi(albumDateTo)
			if err != nil {
				return nil, fmt.Errorf("invalid albumDateTo format: %v", err)
			}
			if (from < 1950 || from > 2025) || (to < 1950 || to > 2025) {
                return nil, fmt.Errorf("invalid albumDateTo format: %v", err)
            }
			if albumYear >= from && albumYear <= to {
				albumDateInRange = true
			} else {
				albumDateInRange = false
			}
		}

		if memberExists && locationExists && creationDateInRange && albumDateInRange {
			filtered = append(filtered, artist)
		}
	}
	return filtered, nil
}
