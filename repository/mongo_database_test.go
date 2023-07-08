package repository

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/wcodesoft/mosha-author-service/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"testing"
)

const databaseName = "mosha"

func createMockedAuthor() bson.D {
	return bson.D{
		{Key: "_id", Value: "ID"},
		{Key: "name", Value: "Name"},
		{Key: "picurl", Value: "PicURL"},
	}
}

func TestMongoDB(t *testing.T) {

	Convey("When using a database instance", t, func() {
		mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
		defer mt.Close()

		mt.Run("Test AddAuthor", func(mt *mtest.T) {
			mt.AddMockResponses(bson.D{{Key: "ok", Value: 1}, {Key: "_id", Value: "ID"}})
			Convey("Test AddAuthor correctly", mt, func() {
				db := NewMongoDatabase(mt.Client, databaseName)
				author := data.Author{ID: "ID", Name: "Name", PicURL: "PicURL"}
				id, err := db.AddAuthor(author)
				So(err, ShouldBeNil)
				So(id, ShouldEqual, author.ID)
			})

			mt.AddMockResponses(bson.D{{Key: "ok", Value: 0}, {Key: "_id", Value: "ID"}})
			Convey("Test AddAuthor with error", mt, func() {
				db := NewMongoDatabase(mt.Client, databaseName)
				author := data.Author{ID: "ID", Name: "Name", PicURL: "PicURL"}
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
				createMockedAuthor(),
			)
			killCursors := mtest.CreateCursorResponse(0, "mosha.authors", mtest.NextBatch)
			mt.AddMockResponses(mockFind, killCursors)
			Convey("Test GetAuthor correctly", mt, func() {
				db := NewMongoDatabase(mt.Client, databaseName)
				author, err := db.GetAuthor("ID")
				So(err, ShouldBeNil)
				So(author.ID, ShouldEqual, "ID")
				So(author.Name, ShouldEqual, "Name")
				So(author.PicURL, ShouldEqual, "PicURL")
			})

			Convey("Test GetAuthor with error", mt, func() {
				db := NewMongoDatabase(mt.Client, databaseName)
				author, err := db.GetAuthor("ID")
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
				err := db.DeleteAuthor("ID")
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
				{Key: "value", Value: createMockedAuthor()}})

			Convey("Test UpdateAuthor correctly", mt, func() {
				db := NewMongoDatabase(mt.Client, databaseName)
				author := data.Author{ID: "ID", Name: "NewName", PicURL: "PicURL"}
				newAuthor, err := db.UpdateAuthor(author)
				So(err, ShouldBeNil)
				So(newAuthor.ID, ShouldEqual, author.ID)
				So(newAuthor.Name, ShouldEqual, author.Name)
				So(newAuthor.PicURL, ShouldEqual, author.PicURL)
			})

			Convey("Test UpdateAuthor with error", mt, func() {
				db := NewMongoDatabase(mt.Client, databaseName)
				author := data.Author{ID: "InvallidID", Name: "NewName", PicURL: "PicURL"}
				newAuthor, err := db.UpdateAuthor(author)
				So(err, ShouldNotBeNil)
				So(newAuthor.ID, ShouldEqual, "")
				So(newAuthor.Name, ShouldEqual, "")
				So(newAuthor.PicURL, ShouldEqual, "")
			})
		})

		mt.Run("Test ListAuthors", func(mt *mtest.T) {
			first := mtest.CreateCursorResponse(
				1,
				"mosha.authors",
				mtest.FirstBatch,
				createMockedAuthor(),
			)
			second := mtest.CreateCursorResponse(
				1,
				"mosha.authors",
				mtest.NextBatch,
				bson.D{
					{Key: "_id", Value: "ID2"},
					{Key: "name", Value: "Name2"},
					{Key: "picurl", Value: "PicURL2"},
				},
			)
			killCursors := mtest.CreateCursorResponse(0, "mosha.authors", mtest.NextBatch)
			mt.AddMockResponses(first, second, killCursors)

			Convey("Test ListAuthors correctly", mt, func() {
				db := NewMongoDatabase(mt.Client, databaseName)
				authors := db.ListAll()
				So(len(authors), ShouldEqual, 2)
			})
		})
	})
}
