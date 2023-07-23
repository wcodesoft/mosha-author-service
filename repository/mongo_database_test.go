package repository

import (
	"testing"

	faker "github.com/brianvoe/gofakeit/v6"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/wcodesoft/mosha-author-service/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

const databaseName = "mosha"

func createMockedAuthor(id string, name string, picUrl string) bson.D {
	return bson.D{
		{Key: "_id", Value: id},
		{Key: "name", Value: name},
		{Key: "picurl", Value: picUrl},
	}
}

func TestMongoDB(t *testing.T) {

	Convey("When using a database instance", t, func() {
		mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
		name := faker.Name()
		picUrl := faker.ImageURL(100, 100)
		id := faker.UUID()
		defer mt.Close()

		mt.Run("Test AddAuthor", func(mt *mtest.T) {
			mt.AddMockResponses(bson.D{{Key: "ok", Value: 1}, {Key: "_id", Value: id}})
			Convey("Test AddAuthor correctly", mt, func() {
				db := NewMongoDatabase(mt.Client, databaseName)
				author := data.Author{ID: id, Name: name, PicURL: picUrl}
				id, err := db.AddAuthor(author)
				So(err, ShouldBeNil)
				So(id, ShouldEqual, author.ID)
			})

			mt.AddMockResponses(bson.D{{Key: "ok", Value: 0}, {Key: "_id", Value: id}})
			Convey("Test AddAuthor with error", mt, func() {
				db := NewMongoDatabase(mt.Client, databaseName)
				author := data.Author{ID: id, Name: name, PicURL: picUrl}
				id, err := db.AddAuthor(author)
				So(err, ShouldNotBeNil)
				So(id, ShouldEqual, "")
			})
		})

		mt.Run("Test GetAuthor", func(mt *mtest.T) {

			mockFind := mtest.CreateCursorResponse(
				1,
				"mosha.authors",
				mtest.FirstBatch,
				createMockedAuthor(id, name, picUrl),
			)
			killCursors := mtest.CreateCursorResponse(0, "mosha.authors", mtest.NextBatch)
			mt.AddMockResponses(mockFind, killCursors)
			Convey("Test GetAuthor correctly", mt, func() {
				db := NewMongoDatabase(mt.Client, databaseName)
				author, err := db.GetAuthor(id)
				So(err, ShouldBeNil)
				So(author.ID, ShouldEqual, id)
				So(author.Name, ShouldEqual, name)
				So(author.PicURL, ShouldEqual, picUrl)
			})

			Convey("Test GetAuthor with error", mt, func() {
				db := NewMongoDatabase(mt.Client, databaseName)
				author, err := db.GetAuthor(id)
				So(err, ShouldNotBeNil)
				So(author.ID, ShouldEqual, "")
				So(author.Name, ShouldEqual, "")
				So(author.PicURL, ShouldEqual, "")
			})
		})

		mt.Run("Test DeleteAuthor", func(mt *mtest.T) {
			Convey("Test DeleteAuthor correctly", mt, func() {
				mt.AddMockResponses(bson.D{{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {Key: "n", Value: 1}})
				db := NewMongoDatabase(mt.Client, databaseName)
				err := db.DeleteAuthor(id)
				So(err, ShouldBeNil)
			})

			Convey("Test DeleteAuthor with error", mt, func() {
				mt.AddMockResponses(bson.D{{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {Key: "n", Value: 0}})
				db := NewMongoDatabase(mt.Client, databaseName)
				err := db.DeleteAuthor("InvalidID")
				So(err, ShouldNotBeNil)
			})
		})

		mt.Run("Test UpdateAuthor", func(mt *mtest.T) {
			mt.AddMockResponses(bson.D{
				{Key: "ok", Value: 1},
				{Key: "value", Value: createMockedAuthor(id, name, picUrl)}})

			Convey("Test UpdateAuthor correctly", mt, func() {
				db := NewMongoDatabase(mt.Client, databaseName)
				newName := faker.Name()
				author := data.Author{ID: id, Name: newName, PicURL: picUrl}
				newAuthor, err := db.UpdateAuthor(author)
				So(err, ShouldBeNil)
				So(newAuthor.ID, ShouldEqual, author.ID)
				So(newAuthor.Name, ShouldEqual, author.Name)
				So(newAuthor.PicURL, ShouldEqual, author.PicURL)
				So(newAuthor.Name, ShouldNotEqual, name)
			})

			Convey("Test UpdateAuthor with error", mt, func() {
				db := NewMongoDatabase(mt.Client, databaseName)
				newName := faker.Name()
				author := data.Author{ID: "InvallidID", Name: newName, PicURL: picUrl}
				newAuthor, err := db.UpdateAuthor(author)
				So(err, ShouldNotBeNil)
				So(newAuthor.ID, ShouldEqual, "")
				So(newAuthor.Name, ShouldEqual, "")
				So(newAuthor.PicURL, ShouldEqual, "")
			})
		})

		mt.Run("Test ListAuthors", func(mt *mtest.T) {
			Convey("Test ListAuthors correctly", mt, func() {
				otherName := faker.Name()
				otherId := faker.UUID()
				first := mtest.CreateCursorResponse(
					1,
					"mosha.authors",
					mtest.FirstBatch,
					createMockedAuthor(id, name, picUrl),
				)
				second := mtest.CreateCursorResponse(
					1,
					"mosha.authors",
					mtest.NextBatch,
					createMockedAuthor(otherId, otherName, picUrl),
				)
				killCursors := mtest.CreateCursorResponse(0, "mosha.authors", mtest.NextBatch)
				mt.AddMockResponses(first, second, killCursors)

				db := NewMongoDatabase(mt.Client, databaseName)
				authors := db.ListAll()
				So(len(authors), ShouldEqual, 2)
			})

			Convey("Test ListAuthors with error", mt, func() {
				mt.AddMockResponses(bson.D{{Key: "ok", Value: 0}})
				db := NewMongoDatabase(mt.Client, databaseName)
				authors := db.ListAll()
				So(len(authors), ShouldNotEqual, 3)
			})
		})
	})
}
