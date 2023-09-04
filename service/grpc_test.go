package service

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/wcodesoft/mosha-author-service/repository"
	pb "github.com/wcodesoft/mosha-service-common/protos/authorservice"
	"google.golang.org/protobuf/types/known/emptypb"

	faker "github.com/brianvoe/gofakeit/v6"
)

func createGrpcRouter() GrpcRouter {
	memoryDatabase := repository.NewInMemoryDatabase()
	clientRepo := repository.NewFakeClientRepository()
	repo := repository.New(memoryDatabase, clientRepo)
	service := New(repo)
	router := NewGrpcRouter(service, "AuthorService")
	return router
}

func TestGrpc(t *testing.T) {
	id := faker.UUID()
	name := faker.Name()
	newName := faker.Name()
	picUrl := faker.ImageURL(100, 100)
	author := &pb.Author{
		Id:     id,
		Name:   name,
		PicUrl: picUrl}
	updatedAuthor := &pb.Author{
		Id:     id,
		Name:   newName,
		PicUrl: picUrl}

	Convey("When adding valid author", t, func() {
		router := createGrpcRouter()
		res, err := router.server.CreateAuthor(context.Background(),
			&pb.CreateAuthorRequest{Author: author},
		)
		Convey("The response should not be nil", func() {
			So(res, ShouldNotBeNil)
		})
		Convey("The error should be nil", func() {
			So(err, ShouldBeNil)
		})
		Convey("The response should contain the correct ID", func() {
			So(res.Id, ShouldEqual, id)
		})
	})

	Convey("With an author in the database", t, func() {
		router := createGrpcRouter()
		router.server.CreateAuthor(context.Background(),
			&pb.CreateAuthorRequest{Author: author},
		)

		Convey("When getting the author", func() {
			res, err := router.server.GetAuthor(context.Background(),
				&pb.GetAuthorRequest{Id: id},
			)
			Convey("The response should not be nil", func() {
				So(res, ShouldNotBeNil)
			})
			Convey("The error should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("The response should contain the correct ID", func() {
				So(res.Id, ShouldEqual, id)
				So(res.Name, ShouldEqual, name)
				So(res.PicUrl, ShouldEqual, picUrl)
			})
		})

		Convey("When deleting the author", func() {
			res, err := router.server.DeleteAuthor(context.Background(),
				&pb.DeleteAuthorRequest{Id: id},
			)
			Convey("The response should not be nil", func() {
				So(res, ShouldNotBeNil)
			})
			Convey("The error should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("The response should contain the correct ID", func() {
				So(res.Success, ShouldEqual, true)
			})
		})

		Convey("When deleting author that does not exist", func() {
			res, err := router.server.DeleteAuthor(context.Background(),
				&pb.DeleteAuthorRequest{Id: faker.UUID()},
			)
			Convey("The response should be nil", func() {
				So(res, ShouldBeNil)
			})
			Convey("The error should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("When updating the author", func() {
			res, err := router.server.UpdateAuthor(context.Background(),
				&pb.UpdateAuthorRequest{Author: updatedAuthor},
			)
			Convey("The response should not be nil", func() {
				So(res, ShouldNotBeNil)
			})
			Convey("The error should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("The response should contain the correct ID", func() {
				So(res.Id, ShouldEqual, id)
				So(res.Name, ShouldEqual, newName)
				So(res.PicUrl, ShouldEqual, picUrl)
			})
		})

		Convey("When listing the authors", func() {
			res, err := router.server.ListAuthors(context.Background(),
				&emptypb.Empty{},
			)
			Convey("The response should not be nil", func() {
				So(res, ShouldNotBeNil)
			})
			Convey("The error should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("The response should contain the correct ID", func() {
				So(res.Authors[0].Id, ShouldEqual, id)
				So(res.Authors[0].Name, ShouldEqual, name)
				So(res.Authors[0].PicUrl, ShouldEqual, picUrl)
			})
		})
	})
}
