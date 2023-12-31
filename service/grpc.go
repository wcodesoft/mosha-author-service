package service

import (
	"context"
	"fmt"
	"github.com/wcodesoft/mosha-author-service/data"
	"github.com/wcodesoft/mosha-service-common/grpc"
	pb "github.com/wcodesoft/mosha-service-common/protos/authorservice"
	"google.golang.org/protobuf/types/known/emptypb"
	"net"
)

// GrpcRouter represents the gRPC router.
type GrpcRouter struct {
	serviceName string
	server      pb.AuthorServiceServer
}

type server struct {
	service Service
	pb.UnimplementedAuthorServiceServer
}

// GetAuthor returns an author by id.
func (g *server) GetAuthor(_ context.Context, request *pb.GetAuthorRequest) (*pb.Author, error) {
	author, err := g.service.GetAuthor(request.Id)
	if err != nil {
		return nil, fmt.Errorf("could not get author: %v", err)
	}
	return toProtoAuthor(author), nil
}

// ListAuthors returns all authors in the database.
func (g *server) ListAuthors(_ context.Context, _ *emptypb.Empty) (*pb.ListAuthorsResponse, error) {
	authors := g.service.ListAll()
	var pbAuthors []*pb.Author
	for _, author := range authors {
		pbAuthors = append(pbAuthors, toProtoAuthor(author))
	}
	return &pb.ListAuthorsResponse{Authors: pbAuthors}, nil
}

// UpdateAuthor updates an author.
func (g *server) UpdateAuthor(_ context.Context, request *pb.UpdateAuthorRequest) (*pb.Author, error) {
	author := request.GetAuthor()
	if author == nil {
		return nil, fmt.Errorf("author is nil")
	}
	updatedAuthor, err := g.service.UpdateAuthor(toAuthorDB(author))
	if err != nil {
		return nil, fmt.Errorf("could not update author: %v", err)
	}
	return toProtoAuthor(updatedAuthor), nil
}

// DeleteAuthor deletes an author by id.
func (g *server) DeleteAuthor(_ context.Context, request *pb.DeleteAuthorRequest) (*pb.DeleteAuthorResponse, error) {
	id := request.GetId()
	err := g.service.DeleteAuthor(id)
	if err != nil {
		return nil, fmt.Errorf("could not delete author: %v", err)
	}
	return &pb.DeleteAuthorResponse{Success: true}, nil
}

// CreateAuthor registers a new Author in the database.
func (g *server) CreateAuthor(_ context.Context, req *pb.CreateAuthorRequest) (*pb.Author, error) {
	author := req.GetAuthor()

	if author == nil {
		return nil, fmt.Errorf("author is nil")
	}

	id, err := g.service.CreateAuthor(toAuthorDB(author))
	if err != nil {
		return nil, err
	}
	return &pb.Author{Id: id, Name: author.Name, PicUrl: author.PicUrl}, nil
}

func toProtoAuthor(author data.Author) *pb.Author {
	return &pb.Author{Id: author.ID, Name: author.Name, PicUrl: author.PicURL}
}

func toAuthorDB(author *pb.Author) data.Author {
	return data.Author{ID: author.Id, Name: author.Name, PicURL: author.PicUrl}
}

func newServer(s Service) pb.AuthorServiceServer {
	return &server{
		service: s,
	}
}

// NewGrpcRouter creates a new gRPC router.
func NewGrpcRouter(s Service, serviceName string) GrpcRouter {
	return GrpcRouter{
		server:      newServer(s),
		serviceName: serviceName,
	}
}

func (g *GrpcRouter) Start(port string) error {
	grpcServer := grpc.CreateNewGRPCServer()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	pb.RegisterAuthorServiceServer(grpcServer, g.server)
	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	return nil
}
