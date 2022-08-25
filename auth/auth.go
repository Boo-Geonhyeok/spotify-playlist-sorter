package auth

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/zmb3/spotify"
)

var redirectURL = "http://127.0.0.1:8080/callback"

// var ClientID = os.Getenv("CLIENT_ID")
// var ClientSecret = os.Getenv("CLIENT_SECERET")
var auth = spotify.NewAuthenticator(redirectURL, spotify.ScopePlaylistReadCollaborative, spotify.ScopePlaylistModifyPublic)

func MakeAuthURL() string {
	err := godotenv.Load()
	if err != nil {
		//send error
	}
	var ClientID = os.Getenv("CLIENT_ID")
	var ClientSecret = os.Getenv("CLIENT_SECERET")
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
