package model

type TrackOption struct {
	PlaylistID   string `json:"playlist_ID"`
	Country      string
	Release_date [2]int //[start, end]
	Genres       []string
	Features     map[string][2]float32
}

func (t *TrackOption) GetPlaylistID() string {
	return t.PlaylistID
}

func (t *TrackOption) GetCountry() string {
	return t.Country
}

func (t *TrackOption) GetReleaseDate() [2]int {
	return t.Release_date
}

func (t *TrackOption) GetGenres() []string {
	return t.Genres
}
