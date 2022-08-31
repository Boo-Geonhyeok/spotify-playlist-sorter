package main

import (
	"net/http"
	"playlist-sorting/handler"
)

func main() {
	http.HandleFunc("/api/url", handler.GetAuthURL)
	http.HandleFunc("/api/extract", handler.ExtractPlaylists)
	http.HandleFunc("/api/genres", handler.GetGenres)
	http.HandleFunc("/api/playlist", handler.GetPlaylists)
	http.HandleFunc("/api/condition", handler.GetCondition)
	http.HandleFunc("/api/filter", handler.GetFilteredPlaylist)

	http.ListenAndServe(":3000", nil)
}
