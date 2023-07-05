package repository

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/wcodesoft/mosha-author-service/data"
	"testing"
)

func TestDatabase(t *testing.T) {
	Convey("When converting author http to database model", t, func() {
		author := authorDB{ID: "ID", Name: "Name", PicURL: "PicURL"}
		authorHttp := toAuthor(author)
		So(authorHttp.ID, ShouldEqual, author.ID)
		So(authorHttp.Name, ShouldEqual, author.Name)
		So(authorHttp.PicURL, ShouldEqual, author.PicURL)
	})

	Convey("When converting author database to http model", t, func() {
		author := data.Author{ID: "ID", Name: "Name", PicURL: "PicURL"}
		authorDb := fromAuthor(author)
		So(authorDb.ID, ShouldEqual, author.ID)
		So(authorDb.Name, ShouldEqual, author.Name)
		So(authorDb.PicURL, ShouldEqual, author.PicURL)
	})
}
