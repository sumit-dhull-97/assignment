package model

type Article struct {
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	UserID    string   `json:"userId"`
	Script    string   `json:"script"`
	Hashtags  []string `json:"hashtags"`
	Created   string   `json:"created"`
	Published string   `json:"published"`
}
