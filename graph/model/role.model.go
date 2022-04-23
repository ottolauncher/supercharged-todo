package model

type Role struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	CreatedAt   *string `json:"createdAt"`
	UpdatedAt   *string `json:"updatedAt"`
	Slug        *string `json:"slug"`
}

func (Role) IsSearchResult() {}
func (Role) IsBaseModel()    {}
