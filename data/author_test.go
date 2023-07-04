package data

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAuthor(t *testing.T) {
	Convey("With AuthorBuilder", t, func() {
		builder := NewAuthorBuilder()

		Convey("When building an author", func() {
			author := builder.Build()

			Convey("The author should be initialized", func() {
				So(author, ShouldNotBeNil)
			})
		})

		Convey("When building an author with a specific ID", func() {
			author := builder.WithId("123").Build()

			Convey("The author should be initialized with the given ID", func() {
				So(author.ID, ShouldEqual, "123")
			})
		})

		Convey("When building an author with a specific name", func() {
			author := builder.WithName("John Doe").Build()

			Convey("The author should be initialized with the given name", func() {
				So(author.Name, ShouldEqual, "John Doe")
			})
		})

		Convey("When building an author with a specific picUrl", func() {
			author := builder.WithPicUrl("http://example.com/pic.jpg").Build()

			Convey("The author should be initialized with the given picUrl", func() {
				So(author.PicURL, ShouldEqual, "http://example.com/pic.jpg")
			})
		})

		Convey("Two authors built with the same builder should be equal", func() {
			author1 := builder.Build()
			author2 := builder.Build()

			Convey("The authors should be equal", func() {
				So(author1, ShouldResemble, author2)
			})
		})
	})
}
