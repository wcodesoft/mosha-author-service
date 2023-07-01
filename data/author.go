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

// AuthorBuilder is the interface that builds an author.
type AuthorBuilder interface {
	WithId(id string) AuthorBuilder
	WithName(name string) AuthorBuilder
	WithPicUrl(picUrl string) AuthorBuilder
	Build() Author
}

type authorBuilder struct {
	id     string
	name   string
	picUrl string
}

// NewAuthorBuilder creates a new author builder.
func NewAuthorBuilder() AuthorBuilder {
	return &authorBuilder{
		id:     uuid.New().String(),
		name:   "",
		picUrl: "",
	}
}

// WithId sets the id of the author.
func (ab *authorBuilder) WithId(id string) AuthorBuilder {
	ab.id = id
	return ab
}

// WithName sets the name of the author.
func (ab *authorBuilder) WithName(name string) AuthorBuilder {
	ab.name = name
	return ab
}

// WithPicUrl sets the picUrl of the author.
func (ab *authorBuilder) WithPicUrl(picUrl string) AuthorBuilder {
	ab.picUrl = picUrl
	return ab
}

// Build builds the author.
func (ab *authorBuilder) Build() Author {
	return Author{
		ID:     ab.id,
		Name:   ab.name,
		PicURL: ab.picUrl,
	}
}
