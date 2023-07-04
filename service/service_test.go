package service

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/wcodesoft/mosha-author-service/data"
	"github.com/wcodesoft/mosha-author-service/repository"
	"testing"
)

func createJohnDoe() data.Author {
	return data.NewAuthorBuilder().
		WithName("John Doe").
		WithPicUrl("http://example.com/john-doe.jpg").
		Build()
}
func TestService(t *testing.T) {

	Convey("When creating a new service", t, func() {
		database := repository.NewInMemoryDatabase()
		repo := repository.New(database)
		service := New(repo)

		Convey("The service should be initialized", func() {
			So(service, ShouldNotBeNil)
		})

		Convey("When adding an author", func() {
			authorId, _ := service.CreateAuthor(createJohnDoe())
			Convey("The list of authors should contain the new author", func() {
				So(len(service.ListAll()), ShouldEqual, 1)
			})

			Convey("Getting the author by ID should return the correct author", func() {
				author, _ := service.GetAuthor(authorId)
				So(author.ID, ShouldEqual, authorId)
				So(author.Name, ShouldEqual, "John Doe")
				So(author.PicURL, ShouldEqual, "http://example.com/john-doe.jpg")
			})
		})

		Convey("When deleting an author", func() {
			authorId, _ := service.CreateAuthor(createJohnDoe())
			err := service.DeleteAuthor(authorId)
			Convey("The list of authors should be empty", func() {
				So(len(service.ListAll()), ShouldEqual, 0)
			})

			Convey("Getting the author by ID should return an error", func() {
				_, getErr := service.GetAuthor(authorId)
				So(getErr, ShouldNotBeNil)
			})

			Convey("Deleting the author should not return an error", func() {
				So(err, ShouldBeNil)
			})
		})

		Convey("When updating an author", func() {
			authorId, _ := service.CreateAuthor(createJohnDoe())

			Convey("Updating the author should return the updated author", func() {
				author, _ := service.UpdateAuthor(data.NewAuthorBuilder().
					WithName("John New Doe").
					WithId(authorId).
					WithPicUrl("http://example.com/john-doe.jpg").
					Build(),
				)
				So(author.ID, ShouldEqual, authorId)
				So(author.Name, ShouldNotEqual, "John Doe")
				So(author.Name, ShouldEqual, "John New Doe")
			})
		})

		Convey("When checking if an author exists", func() {
			authorId, _ := service.CreateAuthor(createJohnDoe())

			Convey("The author should exist", func() {
				So(service.AuthorExist(authorId), ShouldBeTrue)
			})
		})
	})
}
