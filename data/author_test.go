package data

import (
	"testing"

	faker "github.com/brianvoe/gofakeit/v6"
	. "github.com/smartystreets/goconvey/convey"
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
			name := faker.Name()
			author := builder.WithName(name).Build()

			Convey("The author should be initialized with the given name", func() {
				So(author.Name, ShouldEqual, name)
			})
		})

		Convey("When building an author with a specific picUrl", func() {
			picUrl := faker.ImageURL(100, 100)
			author := builder.WithPicUrl(picUrl).Build()

			Convey("The author should be initialized with the given picUrl", func() {
				So(author.PicURL, ShouldEqual, picUrl)
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
