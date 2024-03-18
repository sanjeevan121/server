package routers

import (
	"server/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Music Album Endpoints
	router.HandleFunc("/music_album", handlers.CreateMusicAlbum).Methods("POST")
	router.HandleFunc("/music_album/{id}", handlers.UpdateMusicAlbum).Methods("PUT")

	// Musician Endpoints
	router.HandleFunc("/musician", handlers.CreateMusician).Methods("POST")
	router.HandleFunc("/musician/{id}", handlers.UpdateMusician).Methods("PUT")

	return router
}
