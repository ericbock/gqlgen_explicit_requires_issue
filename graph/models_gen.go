// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

type Pet struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

func (Pet) IsEntity() {}

type Query struct {
}

type User struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Pets     []*Pet   `json:"pets"`
	PetNames []string `json:"petNames"`
}

func (User) IsEntity() {}
