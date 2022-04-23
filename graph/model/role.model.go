package model

import (
	"encoding/json"
	"time"
)

type Role struct {
	ID          string     `json:"id" rethinkdb:"id,omitempty" redis:"id,omitempty"`
	Name        string     `json:"name" rethinkdb:"name" redis:"name"`
	Description *string    `json:"description" rethinkdb:"description" redis:"description"`
	CreatedAt   *time.Time `json:"created_at" rethinkdb:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at" rethinkdb:"updated_at"`
	Slug        *string    `json:"slug" rethinkdb:"slug"`
}

func (Role) IsSearchResult() {}
func (Role) IsBaseModel()    {}

func (r *Role) String() string {
	return r.Name
}

func (r *Role) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, r)
}

func (r *Role) MarshalBinary() ([]byte, error) {
	return json.Marshal(r)
}
