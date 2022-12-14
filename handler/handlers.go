package handler

import (
	"encoding/json"
	"net/http"
	"playlist-sorting/auth"
	"playlist-sorting/filter"
	"playlist-sorting/model"

	"github.com/gocolly/colly/v2"
	"github.com/zmb3/spotify"
)

var client spotify.Client
var playlistIDs []spotify.ID
var playlists []spotify.SimplePlaylist
var trackIDs []spotify.ID
var artists [][]spotify.SimpleArtist
var userTrackOption model.TrackOption
var genreConditions map[string]bool
var countryConditions map[string]bool

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func GetAuthURL(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	url, err := auth.MakeAuthURL()
	if err != nil {
		http.Error(w, err.Error(), 401)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"url": url,
	})
}

func ExtractPlaylists(w http.ResponseWriter, r *http.Request) {
	c := &client
	*c = auth.CreateClient(w, r)
	playlistPage, err := c.CurrentUsersPlaylists()
	playlists = playlistPage.Playlists
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	pi := &playlistIDs
	for _, playlist := range playlists {
		*pi = append(*pi, playlist.ID)
	}
	http.Redirect(w, r, "http://127.0.0.1:8080", http.StatusMovedPermanently)
}

func GetPlaylists(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	json.NewEncoder(w).Encode(playlists)
}

func GetGenres(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	c := colly.NewCollector()
	genres := []string{}
	c.OnHTML("table tbody tr .note a", func(e *colly.HTMLElement) {
		//e.Request.Visit(e.Attr("href"))
		genres = append(genres, e.Text)
	})

	c.Visit("https://everynoise.com/everynoise1d.cgi?scope=all")
	json.NewEncoder(w).Encode(genres)
}

func GetCondition(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&userTrackOption)
	getTracks(userTrackOption.PlaylistID, w)
	countryConditions = map[string]bool{}
	genreConditions = map[string]bool{}
	countryConditions[userTrackOption.Country] = true
	for _, genre := range userTrackOption.Genres {
		genreConditions[genre] = true
	}
}

func GetFilteredPlaylist(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	tt := &trackIDs
	ta := &artists
	if len(userTrackOption.Genres) > 0 {
		*tt, *ta = filter.FilterGenres(w, client, trackIDs, artists, genreConditions)
	}
	if userTrackOption.Release_date != [2]int{0, 0} {
		*tt, *ta = filter.FilterDate(w, client, trackIDs, artists, userTrackOption.Release_date)
	}
	if len(userTrackOption.Features) > 0 {
		*tt, *ta = filter.FilterFeatures(w, client, trackIDs, artists, userTrackOption.Features)
	}
	user, err := client.CurrentUser()
	if err != nil {
		http.Error(w, err.Error(), 401)
	}
	playlist, err := client.CreatePlaylistForUser(user.ID, "sorted", "", true)
	if err != nil {
		http.Error(w, err.Error(), 401)
	}
	_, err = client.AddTracksToPlaylist(playlist.ID, trackIDs...)
	if err != nil {
		http.Error(w, err.Error(), 401)
	}
	json.NewEncoder(w).Encode(playlist)
	reset()
}

func getTracks(userPlaylistID string, w http.ResponseWriter) {
	var playlistID spotify.ID
	//stringPlaylistID := trackOption.GetPlaylistID()
	for _, id := range playlistIDs {
		if id.String() == userPlaylistID {
			playlistID = id
		}
	}
	playlistTrackPage, err := client.GetPlaylistTracks(playlistID)
	if err != nil {
		http.Error(w, err.Error(), 401)
	}

	for _, track := range playlistTrackPage.Tracks {
		trackIDs = append(trackIDs, track.Track.ID)
		artists = append(artists, track.Track.Artists)
	}
}

func reset() {
	playlists = nil
	trackIDs = nil
	artists = nil
}
