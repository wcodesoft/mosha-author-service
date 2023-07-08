package repository

import (
	"github.com/wcodesoft/mosha-author-service/data"
)

type Database interface {
	AddAuthor(author data.Author) (string, error)
	ListAll() []data.Author
	UpdateAuthor(author data.Author) (data.Author, error)
	DeleteAuthor(id string) error
	GetAuthor(id string) (data.Author, error)
}

type authorDB struct {
	ID     string `bson:"_id" json:"id,omitempty"`
	Name   string `bson:"name"`
	PicURL string `bson:"picurl"`
}

func fromAuthor(author data.Author) authorDB {
	return authorDB{
		ID:     author.ID,
		Name:   author.Name,
		PicURL: author.PicURL,
	}
}

func toAuthor(author authorDB) data.Author {
	return data.Author{
		ID:     author.ID,
		Name:   author.Name,
		PicURL: author.PicURL,
	}
}
