package handlers

import (
	"github.com/AlexanderKolesnkov/disign-land/pkg/config"
	"github.com/AlexanderKolesnkov/disign-land/pkg/models"
	"github.com/AlexanderKolesnkov/disign-land/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "In future there will be cool design."

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
