package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Broker struct in charge to handle the servers
type Broker struct {
	config *Config
	router *mux.Router
}

func (b *Broker) Config() *Config {
	return b.config
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	config := b.Config()
	b.router = mux.NewRouter()

	binder(b, b.router)
	// Enabling CORS for every origin
	handler := cors.AllowAll().Handler(b.router)

	/*repo, err := database.NewPostgresRepository(b.config.DatabaseUrl)

	if err != nil {
		log.Fatal("Error connecting to db: ", err)
	}

	repository.SetRepository(repo)*/

	log.Println("Starting server on port", config.Port)

	if err := http.ListenAndServe(config.Port, handler); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
