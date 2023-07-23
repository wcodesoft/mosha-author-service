package service

import (
	"testing"

	faker "github.com/brianvoe/gofakeit/v6"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/wcodesoft/mosha-author-service/data"
	"github.com/wcodesoft/mosha-author-service/repository"
)

func TestService(t *testing.T) {

	name := faker.Name()
	picUrl := faker.ImageURL(100, 100)
	author := data.NewAuthorBuilder().
		WithName(name).
		WithPicUrl(picUrl).
		Build()

	Convey("When creating a new service", t, func() {
		database := repository.NewInMemoryDatabase()
		repo := repository.New(database)
		service := New(repo)

		Convey("The service should be initialized", func() {
			So(service, ShouldNotBeNil)
		})

		Convey("When adding an author", func() {
			authorId, _ := service.CreateAuthor(author)
			Convey("The list of authors should contain the new author", func() {
				So(len(service.ListAll()), ShouldEqual, 1)
			})

			Convey("Getting the author by ID should return the correct author", func() {
				author, _ := service.GetAuthor(authorId)
				So(author.ID, ShouldEqual, authorId)
				So(author.Name, ShouldEqual, name)
				So(author.PicURL, ShouldEqual, picUrl)
			})
		})

		Convey("When deleting an author", func() {
			authorId, _ := service.CreateAuthor(author)
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
			authorId, _ := service.CreateAuthor(author)

			Convey("Updating the author should return the updated author", func() {
				newName := faker.Name()
				author, _ := service.UpdateAuthor(data.NewAuthorBuilder().
					WithId(authorId).
					WithName(newName).
					WithPicUrl(picUrl).
					Build(),
				)
				So(author.ID, ShouldEqual, authorId)
				So(author.Name, ShouldNotEqual, name)
				So(author.Name, ShouldEqual, newName)
			})
		})

	})
}
