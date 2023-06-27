package repository

import (
	"authorservice/data"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
)

type scyllaDatabase struct {
	session     gocqlx.Session
	authorTable *table.Table
}

var authorMetadata = table.Metadata{
	Name:    "authors",
	Columns: []string{"id", "name", "picurl"},
	PartKey: []string{"id"},
}

// AddAuthor adds new author to database.
func (s scyllaDatabase) AddAuthor(author data.Author) (string, error) {
	a := fromAuthor(author)
	q := s.session.Query(s.authorTable.Insert()).BindStruct(a)
	if err := q.ExecRelease(); err != nil {
		return "", err
	}
	return a.ID, nil
}

// ListAll returns all authors in the database.
func (s scyllaDatabase) ListAll() []data.Author {
	var authors []authorDB
	q := s.session.Query(s.authorTable.SelectAll())
	if err := q.SelectRelease(&authors); err != nil {
		panic(err)
	}
	result := make([]data.Author, len(authors))
	for index, author := range authors {
		result[index] = toAuthor(author)
	}
	return result
}

// UpdateAuthor updates an author in the database.
func (s scyllaDatabase) UpdateAuthor(author data.Author) (data.Author, error) {
	a := fromAuthor(author)
	q := s.session.Query(s.authorTable.Update()).BindStruct(a)
	if err := q.ExecRelease(); err != nil {
		return data.Author{}, err
	}
	return toAuthor(a), nil
}

// DeleteAuthor deletes an author from the database.
func (s scyllaDatabase) DeleteAuthor(id string) error {
	q := s.session.Query(s.authorTable.Delete()).BindMap(qb.M{`id`: id})
	if err := q.ExecRelease(); err != nil {
		return err
	}
	return nil
}

// GetAuthor returns an author from the database.
func (s scyllaDatabase) GetAuthor(id string) (data.Author, error) {
	var authors []authorDB
	q := s.session.Query(s.authorTable.Get()).BindMap(qb.M{`id`: id})
	if err := q.SelectRelease(&authors); err != nil {
		return data.Author{}, err
	}
	return toAuthor(authors[0]), nil
}

// AuthorExist checks if an author exists in the database.
func (s scyllaDatabase) AuthorExist(id string) bool {
	var authors []authorDB
	q := s.session.Query(s.authorTable.Get()).BindMap(qb.M{`id`: id})
	if err := q.SelectRelease(&authors); err != nil {
		return false
	}
	return len(authors) > 0
}

// NewScyllaDatabase creates a new Scylla database.
func NewScyllaDatabase(hosts []string, keyspace string) Database {
	cluster := gocql.NewCluster(hosts...)
	cluster.ProtoVersion = 4
	cluster.Keyspace = keyspace

	session, err := gocqlx.WrapSession(cluster.CreateSession())
	prepareDatabase(session, keyspace)

	var authorTable = table.New(authorMetadata)

	if err != nil {
		panic(err)
	}
	return &scyllaDatabase{
		session:     session,
		authorTable: authorTable,
	}
}

func prepareDatabase(session gocqlx.Session, keyspace string) {
	if ok := session.ExecStmt(fmt.Sprintf(
		`CREATE KEYSPACE IF NOT EXISTS  %s WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };`,
		keyspace,
	)); ok != nil {
		panic(ok)
	}

	if ok := session.ExecStmt(fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.authors (
			id text PRIMARY KEY,
			name text,
			picurl text,
	)`, keyspace)); ok != nil {
		panic(ok)
	}
}
