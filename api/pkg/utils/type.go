package utils

type Recommendation struct {
	MovieId  int     `json:"movie_id"`
	Accuracy float64 `json:"accuracy"`
}

type GroupRequest struct {
	GroupName string `json:"name"`
}
