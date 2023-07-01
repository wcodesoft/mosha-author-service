package tests

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/wcodesoft/mosha-author-service/data"
	"testing"
)

func TestAuthor(t *testing.T) {
	Convey("Given a new author with random ID", t, func() {
		author := data.New("John Doe", "http://example.com/john-doe.jpg")

		Convey("The values should be as expected, except for the ID", func() {
			So(author.Name, ShouldEqual, "John Doe")
			So(author.PicURL, ShouldEqual, "http://example.com/john-doe.jpg")
		})
	})

	Convey("Given a new author with a specific ID", t, func() {
		author := data.NewWithId("123", "John Doe", "http://example.com/john-doe.jpg")

		Convey("The values should be as expected, including the ID", func() {
			So(author.ID, ShouldEqual, "123")
			So(author.Name, ShouldEqual, "John Doe")
			So(author.PicURL, ShouldEqual, "http://example.com/john-doe.jpg")
		})
	})
}
