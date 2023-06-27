package tests

import (
	"authorservice/data"
	"authorservice/repository"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRepository(t *testing.T) {

	Convey("Given a new repository", t, func() {
		db := NewInMemoryDatabase()
		repo := repository.New(db)

		Convey("When adding an author", func() {
			id, _ := repo.AddAuthor(data.NewWithId("123", "John Doe", "http://example.com/john-doe.jpg"))

			Convey("The list of authors should contain the new author", func() {
				So(len(repo.ListAll()), ShouldEqual, 1)
			})

			Convey("Adding with same ID should fail", func() {
				_, err := repo.AddAuthor(data.NewWithId(id, "John Doe", "http://example.com/john-doe.jpg"))
				So(err, ShouldNotBeNil)
			})

			Convey("Getting the author by ID should return the correct author", func() {
				author, _ := repo.GetAuthor(id)
				So(author.ID, ShouldEqual, "123")
				So(author.Name, ShouldEqual, "John Doe")
				So(author.PicURL, ShouldEqual, "http://example.com/john-doe.jpg")
			})

			Convey("Updating the author should return the updated author", func() {
				author, _ := repo.UpdateAuthor(data.NewWithId(id, "John New Doe", "http://example.com/john-doe.jpg"))
				So(author.ID, ShouldEqual, "123")
				So(author.Name, ShouldNotEqual, "John Doe")
				So(author.Name, ShouldEqual, "John New Doe")
			})
		})

		Convey("When deleting an author", func() {
			authorID, _ := repo.AddAuthor(data.New("John Doe", "http://example.com/john-doe.jpg"))

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
			_, err := repo.UpdateAuthor(data.New("John Doe", "http://example.com/john-doe.jpg"))
			Convey("An error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("When listing all authors", func() {
			author1, _ := repo.AddAuthor(data.New("John Doe", "http://example.com/john-doe.jpg"))
			author2, _ := repo.AddAuthor(data.New("Jane Doe", "http://example.com/jane-doe.jpg"))

			Convey("The list should contain all authors", func() {
				authors := repo.ListAll()
				So(len(authors), ShouldEqual, 2)
				So(authors[0].ID, ShouldEqual, author1)
				So(authors[1].ID, ShouldEqual, author2)
			})
		})

		Convey("When checking if an author exists", func() {
			authorID, _ := repo.AddAuthor(data.New("John Doe", "http://example.com/john-doe.jpg"))

			Convey("The author should exist", func() {
				So(repo.AuthorExist(authorID), ShouldBeTrue)
			})

			Convey("The author should not exist", func() {
				So(repo.AuthorExist("123"), ShouldBeFalse)
			})
		})
	})
}
