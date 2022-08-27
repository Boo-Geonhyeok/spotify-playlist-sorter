package filter

import (
	"strconv"

	"github.com/zmb3/spotify"
)

func FilterDate(client spotify.Client, trackIDs []spotify.ID, artists [][]spotify.SimpleArtist, dateCondition [2]int) ([]spotify.ID, [][]spotify.SimpleArtist) {
	tmpTrack := trackIDs[:0]
	tmpArtist := artists[:0]

	for index, trackID := range trackIDs {
		fullTrack, err := client.GetTrack(trackID)
		if err != nil {
			//send error
		}
		releaseDate := fullTrack.Album.ReleaseDate[:4]
		c := make(chan bool)
		go matchDate(releaseDate, dateCondition, c)
		x := <-c
		if x == true {
			tmpTrack = append(tmpTrack, trackIDs[index])
			tmpArtist = append(tmpArtist, artists[index])
		}

		// if matchDate(releaseDate, dateCondition) == true {
		// 	tmpTrack = append(tmpTrack, trackIDs[index])
		// 	tmpArtist = append(tmpArtist, artists[index])
		// }
	}
	return tmpTrack, tmpArtist
}

func matchDate(releaseDate string, dateCondition [2]int, c chan bool) {
	date, err := strconv.Atoi(releaseDate)
	if err != nil {
		//send error
	}
	if date >= dateCondition[0] && date <= dateCondition[1] {
		c <- true
		return
		//return true
	}
	c <- false
	return
	//return false
}
