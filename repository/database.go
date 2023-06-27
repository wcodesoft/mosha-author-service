package repository

import (
	"authorservice/data"
)

type Database interface {
	AddAuthor(author data.Author) (string, error)
	ListAll() []data.Author
	UpdateAuthor(author data.Author) (data.Author, error)
	DeleteAuthor(id string) error
	GetAuthor(id string) (data.Author, error)
	AuthorExist(id string) bool
}

type authorDB struct {
	ID     string `db:"id"`
	Name   string `db:"name"`
	PicURL string `db:"picurl"`
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
