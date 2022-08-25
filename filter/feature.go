package filter

import (
	"github.com/zmb3/spotify"
)

var featuresByTrack [][3]float32
var index int

func FilterFeatures(client spotify.Client, trackIDs []spotify.ID, artists [][]spotify.SimpleArtist, featureCondition map[string][2]float32) ([]spotify.ID, [][]spotify.SimpleArtist) {
	tmpTrack := trackIDs[:0]
	tmpArtist := artists[:0]

	audioFeatures, err := client.GetAudioFeatures(trackIDs...)
	if err != nil {
		//send error
	}
	for _, feature := range audioFeatures {
		if _, ok := featureCondition["Danceability"]; ok == false {
			featureCondition["Danceability"] = [2]float32{feature.Danceability, feature.Danceability}
		}
		if _, ok := featureCondition["Instrumentalness"]; ok == false {
			featureCondition["Instrumentalness"] = [2]float32{feature.Instrumentalness, feature.Instrumentalness}
		}
		if _, ok := featureCondition["Valence"]; ok == false {
			featureCondition["Valence"] = [2]float32{feature.Valence, feature.Valence}
		}
		features := [3]float32{feature.Danceability, feature.Instrumentalness, feature.Valence}
		featuresByTrack = append(featuresByTrack, features)
	}

	if matchFeatures(featureCondition) == true {
		tmpTrack = append(tmpTrack, trackIDs[index])
		tmpArtist = append(tmpArtist, artists[index])
	}
	return tmpTrack, tmpArtist
}

func matchFeatures(featureCondition map[string][2]float32) bool {
	for index, features := range featuresByTrack {
		if (features[0] < (featureCondition["Danceability"][0]) || features[0] > (featureCondition["Danceability"][1])) || (features[1] < (featureCondition["Instrumentalness"][0]) || features[1] > (featureCondition["Instrumentalness"][1])) || (features[2] < (featureCondition["Valence"][0]) || features[2] > (featureCondition["Valence"][1])) {
			return false
		}
		i := &index
		*i = index
	}
	return true
}
