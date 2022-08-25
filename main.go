package main

import (
	"net/http"
	"playlist-sorting/handler"
)

func main() {
	http.HandleFunc("/", handler.GetAuthURL)
	http.HandleFunc("/callback", handler.GetPlaylists)
	http.HandleFunc("/condition", handler.GetCondition)
	http.HandleFunc("/filter", handler.GetFilteredPlaylist)

	http.ListenAndServe(":8080", nil)
}
