// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

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

type ArticleInput struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	UserID      string   `json:"userId"`
	Script      string   `json:"script"`
	Hashtags    []string `json:"hashtags"`
	SessionCred string   `json:"sessionCred"`
}

type DeleteInput struct {
	UserID      string `json:"userId"`
	ArticleID   string `json:"articleId"`
	SessionCred string `json:"sessionCred"`
}

type GetAllInput struct {
	UserID      string `json:"userId"`
	SessionCred string `json:"sessionCred"`
}
