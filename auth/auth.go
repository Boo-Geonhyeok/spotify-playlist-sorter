package auth

import (
	"net/http"

	"github.com/zmb3/spotify"
)

var redirectURL = "http://127.0.0.1:8080/callback"
var ClientID = "70d3fbedecc646b184fb5407783140dd"
var ClientSecret = "9781d25cd84041ae802c764d0694b6a7"
var auth = spotify.NewAuthenticator(redirectURL, spotify.ScopePlaylistReadCollaborative, spotify.ScopePlaylistModifyPublic)

func MakeAuthURL() string {
	auth.SetAuthInfo(ClientID, ClientSecret)
	url := auth.AuthURL("state")
	return url
}

func CreateClient(w http.ResponseWriter, r *http.Request) (client spotify.Client) {
	token, err := auth.Token("state", r)

	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusNotFound)
		return
	}

	client = auth.NewClient(token)
	return client
}
