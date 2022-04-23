package model

import (
	"encoding/json"
	"time"
)

type Todo struct {
	ID          string     `json:"id" rethinkdb:"id,omitempty" redis:"id,omitempty"`
	Description *string    `json:"description" rethinkdb:"description" redis:"description"`
	Text        string     `json:"text" rethinkdb:"text" redis:"text"`
	Done        bool       `json:"done" rethinkdb:"done" redis:"done"`
	User        *User      `json:"user" rethinkdb:"user" redis:"user"`
	UserID      string     `json:"user_id" rethinkdb:"user_id,omitempty" redis:"user_id,omitempty"`
	Assigned    []*User    `json:"assigned" rethinkdb:"assigned" redis:"assigned"`
	AssignedIDs []string   `json:"assigned_ids" rethinkdb:"assigned_ids,omitempty" redis:"assigned_ids,omitempty"`
	CreatedAt   *time.Time `json:"created_at" rethinkdb:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at" rethinkdb:"updated_at"`
	Slug        *string    `json:"slug" rethinkdb:"slug,omitempty"`
	Start       time.Time  `json:"start" rethinkdb:"start"`
	End         *time.Time `json:"end" rethinkdb:"end"`
}

func (Todo) IsSearchResult() {}
func (Todo) IsBaseModel()    {}

func (t *Todo) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, t)
}

func (t *Todo) MarshalBinary() ([]byte, error) {
	return json.Marshal(t)
}
