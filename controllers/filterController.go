package controllers

import (
	"groupie_tracker/models"
	"strconv"
	"time"
)

func Filter(artists []models.Artist, members string,  location string, creationDateFrom string, creationDateTo string, albumDateFrom string, albumDateTo string) ([]models.Artist, error) {
	var filtered []models.Artist
	var memberExists bool
	var locationExists bool
	var creationDateFromExists bool
	var creationDateToExists bool
	var albumDateFromExists bool
	var albumDateToExists bool
	for _, artist := range artists {
		if len(members) > 0 {
			for _, member := range members{
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
			if locations == location {
				locationExists = true
			}
		}
		}
		if len(creationDateFrom) > 0 {
			from, err := strconv.Atoi( creationDateFrom)
			if err != nil {
				return nil, err
			}
			if artist.CreationDate > from {
				creationDateFromExists = true
			}
			
		
		}
		if len(creationDateTo) > 0 {
			to, err := strconv.Atoi(creationDateTo)
			if err != nil {
				return nil, err
			}
			if artist.CreationDate < to {
				creationDateToExists = true
				}
		
		}
		if len(albumDateFrom) > 0 {
			from, err := time.Parse("2006-01-02", albumDateFrom)
			if err!= nil {
                return nil, err
            }
			FirstAlbum, err := time.Parse("2006-01-02", artist.FirstAlbum)
			if err != nil {
				return nil, err
			}
			if FirstAlbum.After(from) {
				albumDateFromExists = true
			}
		
		}
		if len(albumDateTo) > 0 {
			to, err := time.Parse("2006-01-02", albumDateTo)
			if err != nil {
				return nil, err
			}
			FirstAlbum, err := time.Parse("2006-01-02", artist.FirstAlbum)
			if err != nil {
				return nil, err
			}
			if FirstAlbum.Before(to) {
				albumDateToExists = true
			}
		
		}
		if memberExists || locationExists || creationDateFromExists || creationDateToExists || albumDateFromExists || albumDateToExists {
			filtered = append(filtered, artist)
		}
		}
	return filtered, nil
}
