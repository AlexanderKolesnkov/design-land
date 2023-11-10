package main

import (
	"fmt"
	"github.com/AlexanderKolesnkov/disign-land/pkg/config"
	"github.com/AlexanderKolesnkov/disign-land/pkg/handlers"
	"github.com/AlexanderKolesnkov/disign-land/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalln("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Repo.Home)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
