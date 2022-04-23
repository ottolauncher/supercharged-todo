package model

import (
	"encoding/json"
	"time"
)

type User struct {
	ID        string     `json:"id" rethinkdb:"id,omitempty" redis:"id,omitempty"`
	Username  string     `json:"username" rethinkdb:"username" redis:"username"`
	Slug      *string    `json:"slug" rethinkdb:"slug"`
	Email     string     `json:"email" rethinkdb:"email" redis:"email"`
	Password  string     `json:"password" rethinkdb:"password" redis:"password"`
	Biography *string    `json:"biography" rethinkdb:"biography,omitempty" redis:"biography,omitempty"`
	Roles     []*Role    `json:"roles" rethinkdb:"roles" redis:"roles"`
	RoleIDs   []string   `json:"role_ids" rethinkdb:"role_ids" redis:"role_ids"`
	LastLogin *time.Time `json:"last_login" rethinkdb:"last_login"`
	CreatedAt *time.Time `json:"created_at" rethinkdb:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" rethinkdb:"updated_at"`
}

func (User) IsSearchResult() {}
func (User) IsBaseModel()    {}

func (u *User) String() string {
	return u.Username
}

func (u *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}

func (u *User) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}
