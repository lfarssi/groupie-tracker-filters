package controllers

import (
	"groupie_tracker/models"
	"strconv"
	"time"
)

func filter(artists []models.Artist, members string,  location string, creationDateFrom string, creationDateTo string, albumDateFrom string, albumDateTo string) ([]models.Artist, error) {
	var filtered []models.Artist
	for _, artist := range artists {
		if len(members) > 0 {
			members, err := strconv.Atoi(members)
			if err != nil {
				return nil, err
			}
			if len(artist.Members) != members {
				continue
			}
		}
		if len(location) > 0 {
			for _, location := range artist.Locationsr {
			if location != location {
				continue
			}
		}
		}
		if len(creationDateFrom) > 0 {
			from, err := strconv.Atoi( creationDateFrom)
			if err != nil {
				return nil, err
			}
			if artist.CreationDate < from {
					continue
			}
			
		
		}
		if len(creationDateTo) > 0 {
			to, err := strconv.Atoi( creationDateFrom)
			if err != nil {
				return nil, err
			}
			if artist.CreationDate > to {
					continue
				}
		
		}
		if len(albumDateFrom) > 0 {
			from, err := time.Parse("2006-01-02", albumDateFrom)
			FirstAlbum, err := time.Parse("2006-01-02", artist.FirstAlbum)

			if err != nil {
				return nil, err
			}
			if FirstAlbum.Before(from) {
				continue
			}
		
		}
		if len(albumDateTo) > 0 {
			to, err := time.Parse("2006-01-02", albumDateTo)
			FirstAlbum, err := time.Parse("2006-01-02", artist.FirstAlbum)
			if err != nil {
				return nil, err
			}
			if FirstAlbum.After(to) {
				continue
			}
		
		}
		filtered = append(filtered, artist)
	}
	return filtered, nil
}
