package filter

import (
	"github.com/zmb3/spotify"
)

func FilterGenres(client spotify.Client, trackIDs []spotify.ID, artists [][]spotify.SimpleArtist, genreConditions map[string]bool) ([]spotify.ID, [][]spotify.SimpleArtist) {
	tmpTrack := trackIDs[:0]
	tmpArtist := artists[:0]

	for index, artist := range artists {
		for _, a := range artist {
			fullArtist, err := client.GetArtist(a.ID)
			if err != nil {
				//todo: send error
			}
			c := make(chan bool)
			go matchGenres(fullArtist.Genres, genreConditions, c)
			x := <-c
			if x == true {
				tmpTrack = append(tmpTrack, trackIDs[index])
				tmpArtist = append(tmpArtist, artists[index])
				break
			}
			// if matchGenres(fullArtist.Genres, genreConditions) == true {
			// 	// trackIDs = append(trackIDs[:index], trackIDs[index+1:]...)
			// 	tmpTrack = append(tmpTrack, trackIDs[index])
			// 	tmpArtist = append(tmpArtist, artists[index])
			// 	break
			// }
		}
	}
	return tmpTrack, tmpArtist
}

func matchGenres(artistGenres []string, genreConditions map[string]bool, c chan bool) {
	for _, genre := range artistGenres {
		if genreConditions[genre] {
			//return true
			c <- true
			return
		}
	}
	c <- false
	return
	//return false
}
