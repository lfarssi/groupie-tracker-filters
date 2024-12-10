package routes

import (
	"groupie_tracker/controllers"
	"log"
	"net/http"
)

func Router() {
	http.HandleFunc("/", controllers.ArtistController)
	http.HandleFunc("/static/", controllers.StaticController)
	http.HandleFunc("/artist/{id}", controllers.ArtistDetailController)
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
