package model

type Todo struct {
	ID          string   `json:"id"`
	Text        string   `json:"text"`
	Done        bool     `json:"done"`
	User        *User    `json:"user"`
	UserID      string   `json:"userID"`
	Assigned    []*User  `json:"assigned"`
	AssignedIDs []string `json:"assignedIDs"`
	Description *string  `json:"description"`
	Start       string   `json:"start"`
	End         *string  `json:"end"`
	CreatedAt   *string  `json:"createdAt"`
	UpdatedAt   *string  `json:"updatedAt"`
	Slug        *string  `json:"slug"`
}

func (Todo) IsSearchResult() {}
func (Todo) IsBaseModel()    {}
