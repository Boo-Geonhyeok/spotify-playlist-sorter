package handler

import (
	"encoding/json"
	"net/http"
	"playlist-sorting/auth"
	"playlist-sorting/filter"
	"playlist-sorting/model"

	"github.com/zmb3/spotify"
)

var client spotify.Client
var playlistIDs []spotify.ID
var trackIDs []spotify.ID
var artists [][]spotify.SimpleArtist
var userTrackOption model.TrackOption
var genreConditions map[string]bool
var countryConditions map[string]bool

func GetAuthURL(w http.ResponseWriter, r *http.Request) {
	url, err := auth.MakeAuthURL()
	if err != nil {
		http.Error(w, err.Error(), 401)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"url": url,
	})
}

func GetPlaylists(w http.ResponseWriter, r *http.Request) {
	c := &client
	*c = auth.CreateClient(w, r)
	playlistPage, err := c.CurrentUsersPlaylists()
	playlists := playlistPage.Playlists
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	pi := &playlistIDs
	for _, playlist := range playlists {
		*pi = append(*pi, playlist.ID)
	}
	json.NewEncoder(w).Encode(playlists)
}

func GetCondition(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&userTrackOption)
	getTracks(userTrackOption, w)
	countryConditions = map[string]bool{}
	genreConditions = map[string]bool{}
	countryConditions[userTrackOption.Country] = true
	for _, genre := range userTrackOption.Genres {
		genreConditions[genre] = true
	}
}

func GetFilteredPlaylist(w http.ResponseWriter, r *http.Request) {
	tt := &trackIDs
	ta := &artists
	if userTrackOption.Genres != nil {
		*tt, *ta = filter.FilterGenres(w, client, trackIDs, artists, genreConditions)
	}
	if userTrackOption.Release_date != [2]int{0, 0} {
		*tt, *ta = filter.FilterDate(w, client, trackIDs, artists, userTrackOption.Release_date)
	}
	if userTrackOption.Features != nil {
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
}

func getTracks(trackOption model.TrackOption, w http.ResponseWriter) {
	var playlistID spotify.ID
	stringPlaylistID := trackOption.GetPlaylistID()
	for _, id := range playlistIDs {
		if id.String() == stringPlaylistID {
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
