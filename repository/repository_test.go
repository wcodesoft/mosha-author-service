package repository

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/wcodesoft/mosha-author-service/data"
	"testing"
)

func TestRepository(t *testing.T) {

	Convey("Given a new repository", t, func() {
		db := NewInMemoryDatabase()
		repo := New(db)

		Convey("When adding an author", func() {
			builder := data.
				NewAuthorBuilder().
				WithId("123").
				WithName("John Doe").
				WithPicUrl("http://example.com/john-doe.jpg")
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
				author, _ := repo.GetAuthor(id)
				So(author.ID, ShouldEqual, "123")
				So(author.Name, ShouldEqual, "John Doe")
				So(author.PicURL, ShouldEqual, "http://example.com/john-doe.jpg")
			})

			Convey("Updating the author should return the updated author", func() {
				author, _ := repo.UpdateAuthor(
					builder.
						WithId(id).
						WithName("John New Doe").
						WithPicUrl("http://example.com/john-doe.jpg").
						Build(),
				)
				So(author.ID, ShouldEqual, "123")
				So(author.Name, ShouldNotEqual, "John Doe")
				So(author.Name, ShouldEqual, "John New Doe")
			})
		})

		Convey("When deleting an author", func() {
			authorID, _ := repo.AddAuthor(data.NewAuthorBuilder().WithName("John Doe").Build())

			Convey("Deleting the author should remove it from the list", func() {
				if err := repo.DeleteAuthor(authorID); err != nil {
					t.Fatal(err)
				}
				So(len(repo.ListAll()), ShouldEqual, 0)
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
				WithName("John Doe").
				WithPicUrl("http://example.com/john-doe.jpg").
				Build())
			_, _ = repo.AddAuthor(
				data.NewAuthorBuilder().
					WithName("Jane Doe").
					WithPicUrl("http://example.com/john-doe.jpg").
					Build())

			Convey("The list should contain all authors", func() {
				authors := repo.ListAll()
				So(len(authors), ShouldEqual, 2)
			})
		})
	})
}
