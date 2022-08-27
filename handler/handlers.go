package handler

import (
	"encoding/json"
	"fmt"
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
	url := auth.MakeAuthURL()
	fmt.Fprintln(w, url)
	//todo: convert url to json format
}

func GetPlaylists(w http.ResponseWriter, r *http.Request) {
	c := &client
	*c = auth.CreateClient(w, r)
	playlistPage, err := c.CurrentUsersPlaylists()
	playlists := playlistPage.Playlists
	if err != nil {
		//todo: send error
		return
	}
	pi := &playlistIDs
	for _, playlist := range playlists {
		*pi = append(*pi, playlist.ID)
	}
	fmt.Fprintln(w, playlists)
	//todo: convert playlists to json format
}

func GetCondition(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&userTrackOption)
	getTracks(userTrackOption)
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
		*tt, *ta = filter.FilterGenres(client, trackIDs, artists, genreConditions)
	}
	if userTrackOption.Release_date != [2]int{0, 0} {
		*tt, *ta = filter.FilterDate(client, trackIDs, artists, userTrackOption.Release_date)
	}
	if userTrackOption.Features != nil {
		*tt, *ta = filter.FilterFeatures(client, trackIDs, artists, userTrackOption.Features)
	}
	fmt.Println(trackIDs)
	// user, err := client.CurrentUser()
	// if err != nil {
	// 	//send error
	// }
	// playlist, err := client.CreatePlaylistForUser(user.ID, "sorted", "", true)
	// if err != nil {
	// 	//send error
	// }
	// _, err = client.AddTracksToPlaylist(playlist.ID, trackIDs...)
	// if err != nil {
	// 	//send error
	// }
	// //send playlist
}

func getTracks(trackOption model.TrackOption) {
	var playlistID spotify.ID
	stringPlaylistID := trackOption.GetPlaylistID()
	for _, id := range playlistIDs {
		if id.String() == stringPlaylistID {
			playlistID = id
		}
	}
	playlistTrackPage, err := client.GetPlaylistTracks(playlistID)
	if err != nil {
		//todo: send error
	}

	for _, track := range playlistTrackPage.Tracks {
		trackIDs = append(trackIDs, track.Track.ID)
		artists = append(artists, track.Track.Artists)
	}
}
