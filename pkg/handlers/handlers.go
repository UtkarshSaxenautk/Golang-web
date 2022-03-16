package handlers

import (
	"github.com/utkarshsaxenautk/pkg/config"
	"github.com/utkarshsaxenautk/pkg/models"
	"net/http"

	"github.com/utkarshsaxenautk/pkg/render"
)

// Hold data for handlers

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// Creates new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Sets repository for handler
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	intMap := make(map[string]int)
	stringMap["test2"] = "Hello Walker"
	intMap["myval"] = 67
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		IntMap:    intMap,
		StringMap: stringMap,
	})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Utk"
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}
