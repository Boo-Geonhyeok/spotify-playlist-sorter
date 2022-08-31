package filter

import (
	"net/http"

	"github.com/zmb3/spotify"
)

var featuresByTrack [][3]float32
var featureConditionByTrack []map[string][2]float32
var indexing int
var userFeatureCondition map[string][2]float32

func FilterFeatures(w http.ResponseWriter, client spotify.Client, trackIDs []spotify.ID, artists [][]spotify.SimpleArtist, featureCondition map[string][2]float32) ([]spotify.ID, [][]spotify.SimpleArtist) {
	tmpTrack := trackIDs[:0]
	tmpArtist := artists[:0]
	audioFeatures, err := client.GetAudioFeatures(trackIDs...)
	if err != nil {
		http.Error(w, err.Error(), 401)
	}
	userFeatureCondition = featureCondition
	for _, feature := range audioFeatures {
		checkFeatures(feature, featureCondition)
	}

	c := make(chan int)
	go matchFeatures(featureCondition, c)
	for i := range c {
		tmpTrack = append(tmpTrack, trackIDs[i])
		tmpArtist = append(tmpArtist, artists[indexing])
	}
	return tmpTrack, tmpArtist
}

func matchFeatures(featureCondition map[string][2]float32, c chan int) {
	for index, features := range featuresByTrack {
		if features[0] < (featureConditionByTrack[index]["Danceability"][0]) || features[0] > (featureConditionByTrack[index]["Danceability"][1]) {
			continue
		} else if features[1] < (featureConditionByTrack[index]["Instrumentalness"][0]) || features[1] > (featureConditionByTrack[index]["Instrumentalness"][1]) {
			continue
		} else if features[2] < (featureConditionByTrack[index]["Valence"][0]) || features[2] > (featureConditionByTrack[index]["Valence"][1]) {
			continue
		}
		c <- index
	}
	close(c)
}

func checkFeatures(feature *spotify.AudioFeatures, featureCondition map[string][2]float32) {
	f := map[string][2]float32{}
	if _, ok := userFeatureCondition["Danceability"]; ok == false {
		//featureCondition["Danceability"] = [2]float32{feature.Danceability, feature.Danceability}
		f["Danceability"] = [2]float32{feature.Danceability, feature.Danceability}
	} else {
		f["Danceability"] = userFeatureCondition["Danceability"]
	}

	if _, ok := userFeatureCondition["Instrumentalness"]; ok == false {
		//featureCondition["Instrumentalness"] = [2]float32{feature.Instrumentalness, feature.Instrumentalness}
		f["Instrumentalness"] = [2]float32{feature.Instrumentalness, feature.Instrumentalness}
	} else {
		f["Instrumentalness"] = userFeatureCondition["Instrumentalness"]
	}

	if _, ok := userFeatureCondition["Valence"]; ok == false {
		//featureCondition["Valence"] = [2]float32{feature.Valence, feature.Valence}
		f["Valence"] = [2]float32{feature.Valence, feature.Valence}
	} else {
		f["Valence"] = userFeatureCondition["Valence"]
	}

	features := [3]float32{feature.Danceability, feature.Instrumentalness, feature.Valence}
	featuresByTrack = append(featuresByTrack, features)
	featureConditionByTrack = append(featureConditionByTrack, f)
}
