package filter

import (
	"net/http"

	"github.com/zmb3/spotify"
)

func FilterGenres(w http.ResponseWriter, client spotify.Client, trackIDs []spotify.ID, artists [][]spotify.SimpleArtist, genreConditions map[string]bool) ([]spotify.ID, [][]spotify.SimpleArtist) {
	tmpTrack := trackIDs[:0]
	tmpArtist := artists[:0]

	for index, artist := range artists {
		for _, a := range artist {
			fullArtist, err := client.GetArtist(a.ID)
			if err != nil {
				http.Error(w, err.Error(), 401)
			}
			c := make(chan bool)
			go matchGenres(fullArtist.Genres, genreConditions, c)
			x := <-c
			if x == true {
				tmpTrack = append(tmpTrack, trackIDs[index])
				tmpArtist = append(tmpArtist, artists[index])
				break
			}
		}
	}
	return tmpTrack, tmpArtist
}

func matchGenres(artistGenres []string, genreConditions map[string]bool, c chan bool) {
	for _, genre := range artistGenres {
		if genreConditions[genre] {
			c <- true
			return
		}
	}
	c <- false
	return
}
