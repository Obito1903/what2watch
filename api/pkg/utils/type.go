package utils

type Recommendation struct {
	MovieId  int     `json:"movie_id"`
	Accuracy float64 `json:"accuracy"`
}

type Group struct {
	GroupName string `json:"name"`
	UserID    int    `json:"group_id"`
}
