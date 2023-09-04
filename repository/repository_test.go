package repository

import (
	"fmt"
	"testing"

	faker "github.com/brianvoe/gofakeit/v6"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/wcodesoft/mosha-author-service/data"
)

func TestRepository(t *testing.T) {

	Convey("Given a new repository", t, func() {
		db := NewInMemoryDatabase()
		clientRepository := NewFakeClientRepository()
		repo := New(db, clientRepository)
		fakeId := faker.UUID()
		name := faker.Name()
		picUrl := faker.ImageURL(100, 100)

		Convey("When adding an author", func() {
			builder := data.
				NewAuthorBuilder().
				WithId(fakeId).
				WithName(name).
				WithPicUrl(picUrl)
			author := builder.Build()
			id, _ := repo.AddAuthor(author)

			Convey("The list of authors should contain the new author", func() {
				So(len(repo.ListAll()), ShouldEqual, 1)
			})

			Convey("Adding with same ID should fail", func() {
				_, err := repo.AddAuthor(author)
				So(err, ShouldNotBeNil)
			})

			Convey("Getting the author by ID should return the correct author", func() {
				author, _ := repo.GetAuthor(fakeId)
				So(author.ID, ShouldEqual, fakeId)
				So(author.Name, ShouldEqual, name)
				So(author.PicURL, ShouldEqual, picUrl)
			})

			Convey("Updating the author should return the updated author", func() {
				newName := faker.Name()
				author, _ := repo.UpdateAuthor(
					builder.
						WithId(id).
						WithName(newName).
						WithPicUrl(picUrl).
						Build(),
				)
				So(author.ID, ShouldEqual, id)
				So(author.Name, ShouldNotEqual, name)
				So(author.Name, ShouldEqual, newName)
			})
		})

		Convey("When deleting an author", func() {
			authorID, _ := repo.AddAuthor(data.NewAuthorBuilder().WithName(name).Build())

			Convey("Deleting the author should remove it from the list", func() {
				if err := repo.DeleteAuthor(authorID); err != nil {
					t.Fatal(err)
				}
				So(len(repo.ListAll()), ShouldEqual, 0)
			})
		})

		Convey("When deleting an author with errors on quote service", func() {
			authorID, _ := repo.AddAuthor(data.NewAuthorBuilder().WithName(name).Build())

			Convey("When quotes service throw error, should not delete author", func() {
				clientRepository.SetDeleteAuthorQuotesReturn(fmt.Errorf("error"))
				err := repo.DeleteAuthor(authorID)
				So(err, ShouldNotBeNil)
				So(len(repo.ListAll()), ShouldEqual, 1)
			})
		})

		Convey("When deleting an author that does not exist", func() {
			err := repo.DeleteAuthor("123")
			Convey("An error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("When getting an author that does not exist", func() {
			_, err := repo.GetAuthor("123")
			Convey("An error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("When updating an author that does not exist", func() {
			_, err := repo.UpdateAuthor(data.NewAuthorBuilder().Build())
			Convey("An error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("When listing all authors", func() {
			_, _ = repo.AddAuthor(data.NewAuthorBuilder().
				WithName(name).
				WithPicUrl(picUrl).
				Build())
			_, _ = repo.AddAuthor(
				data.NewAuthorBuilder().
					WithName(faker.Name()).
					WithPicUrl(faker.ImageURL(100, 100)).
					Build())

			Convey("The list should contain all authors", func() {
				authors := repo.ListAll()
				So(len(authors), ShouldEqual, 2)
			})
		})
	})
}
