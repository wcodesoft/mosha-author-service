package service

import (
	"context"
	. "github.com/smartystreets/goconvey/convey"
	pb "github.com/wcodesoft/mosha-author-service/proto"
	"github.com/wcodesoft/mosha-author-service/repository"
	"google.golang.org/protobuf/types/known/emptypb"
	"testing"
)

func createGrpcRouter() GrpcRouter {
	memoryDatabase := repository.NewInMemoryDatabase()
	repo := repository.New(memoryDatabase)
	service := New(repo)
	router := NewGrpcRouter(service, "AuthorService")
	return router
}

func TestGrpc(t *testing.T) {
	Convey("When adding valid author", t, func() {
		router := createGrpcRouter()
		res, err := router.server.CreateAuthor(context.Background(),
			&pb.CreateAuthorRequest{Author: &pb.Author{
				Id:     "123",
				Name:   "John Doe",
				PicUrl: "PicUrl"},
			},
		)
		Convey("The response should not be nil", func() {
			So(res, ShouldNotBeNil)
		})
		Convey("The error should be nil", func() {
			So(err, ShouldBeNil)
		})
		Convey("The response should contain the correct ID", func() {
			So(res.Id, ShouldEqual, "123")
		})
	})

	Convey("With an author in the database", t, func() {
		router := createGrpcRouter()
		router.server.CreateAuthor(context.Background(),
			&pb.CreateAuthorRequest{Author: &pb.Author{
				Id:     "123",
				Name:   "John Doe",
				PicUrl: "PicUrl"},
			},
		)

		Convey("When getting the author", func() {
			res, err := router.server.GetAuthor(context.Background(),
				&pb.GetAuthorRequest{Id: "123"},
			)
			Convey("The response should not be nil", func() {
				So(res, ShouldNotBeNil)
			})
			Convey("The error should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("The response should contain the correct ID", func() {
				So(res.Id, ShouldEqual, "123")
				So(res.Name, ShouldEqual, "John Doe")
				So(res.PicUrl, ShouldEqual, "PicUrl")
			})
		})

		Convey("When deleting the author", func() {
			res, err := router.server.DeleteAuthor(context.Background(),
				&pb.DeleteAuthorRequest{Id: "123"},
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
				&pb.DeleteAuthorRequest{Id: "426"},
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
				&pb.UpdateAuthorRequest{Author: &pb.Author{
					Id:     "123",
					Name:   "Jane Doe",
					PicUrl: "PicUrl"},
				},
			)
			Convey("The response should not be nil", func() {
				So(res, ShouldNotBeNil)
			})
			Convey("The error should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("The response should contain the correct ID", func() {
				So(res.Id, ShouldEqual, "123")
				So(res.Name, ShouldEqual, "Jane Doe")
				So(res.PicUrl, ShouldEqual, "PicUrl")
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
				So(res.Authors[0].Id, ShouldEqual, "123")
				So(res.Authors[0].Name, ShouldEqual, "John Doe")
				So(res.Authors[0].PicUrl, ShouldEqual, "PicUrl")
			})
		})
	})
}
