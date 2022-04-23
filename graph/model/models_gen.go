// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type BaseModel interface {
	IsBaseModel()
}

type SearchResult interface {
	IsSearchResult()
}

type NewRole struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type NewTodo struct {
	Text        string     `json:"text"`
	UserID      string     `json:"userId"`
	UserIds     []string   `json:"userIds"`
	Description *string    `json:"description"`
	End         *time.Time `json:"end"`
	Start       time.Time  `json:"start"`
}

type NewUser struct {
	Username  string   `json:"username"`
	Password1 string   `json:"password1"`
	Password2 string   `json:"password2"`
	Email     string   `json:"email"`
	Biography *string  `json:"biography"`
	Roles     []string `json:"roles"`
}

type UpdateRole struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type UpdateTodo struct {
	ID          string     `json:"id"`
	Text        string     `json:"text"`
	UserID      string     `json:"userId"`
	UserIds     []string   `json:"userIds"`
	Description *string    `json:"description"`
	End         *time.Time `json:"end"`
	Start       time.Time  `json:"start"`
}

type UpdateUser struct {
	ID          string   `json:"id"`
	Username    string   `json:"username"`
	OldPassword string   `json:"oldPassword"`
	Password1   string   `json:"password1"`
	Password2   string   `json:"password2"`
	Email       string   `json:"email"`
	Biography   *string  `json:"biography"`
	Roles       []string `json:"roles"`
}
