package handlers

import (
	_ "fmt"
	"github.com/DianaSun97/elluliin_booking/render"
	_ "html/template"
	"net/http"
)

// Home is the handler for the home page
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home-page.gohtml")
}

// About is the handler for the about page
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about-page.gohtml")
}
