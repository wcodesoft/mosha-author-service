package data

import "github.com/google/uuid"

// Author represents an author.
type Author struct {
	// ID is the unique identifier of the author.
	ID string `json:"id"`
	//	Name is the name of the author.
	Name string `json:"name"`
	//	PicURL is the URL of the author's picture.
	PicURL string `json:"picUrl"`
}

// NewWithId creates a new author with the given id, name and picUrl.
func NewWithId(id, name, picUrl string) Author {
	return Author{
		ID:     id,
		Name:   name,
		PicURL: picUrl,
	}
}

func New(name, picUrl string) Author {
	return Author{
		ID:     uuid.New().String(),
		Name:   name,
		PicURL: picUrl,
	}
}
