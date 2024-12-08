package controllers

import (
	"html/template"
	"log"
	"net/http"
)

func ParseController( w http.ResponseWriter, r *http.Request,tmpl string, data any) {
	res, err := template.ParseFiles("views/"+tmpl+".html")
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError)
		return
	}

	if err = res.Execute(w, data); err != nil {
		log.Println(err)
		ErrorController(w, r, http.StatusInternalServerError)
		return
	}
}
