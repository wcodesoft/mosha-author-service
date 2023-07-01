package service

import (
	"authorservice/data"
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "authorservice/proto"
)

// GrpcRouter represents the gRPC router.
type GrpcRouter struct {
	service Service
	pb.UnimplementedAuthorServiceServer
}

// GetAuthor returns an author by id.
func (g GrpcRouter) GetAuthor(_ context.Context, request *pb.GetAuthorRequest) (*pb.Author, error) {
	author, err := g.service.GetAuthor(request.Id)
	if err != nil {
		return nil, fmt.Errorf("could not get author: %v", err)
	}
	return toProtoAuthor(author), nil
}

// ListAuthors returns all authors in the database.
func (g GrpcRouter) ListAuthors(_ context.Context, _ *emptypb.Empty) (*pb.ListAuthorsResponse, error) {
	authors := g.service.ListAll()
	var pbAuthors []*pb.Author
	for _, author := range authors {
		pbAuthors = append(pbAuthors, toProtoAuthor(author))
	}
	return &pb.ListAuthorsResponse{Authors: pbAuthors}, nil
}

// UpdateAuthor updates an author.
func (g GrpcRouter) UpdateAuthor(_ context.Context, request *pb.UpdateAuthorRequest) (*pb.Author, error) {
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
func (g GrpcRouter) DeleteAuthor(_ context.Context, request *pb.DeleteAuthorRequest) (*pb.DeleteAuthorResponse, error) {
	id := request.GetId()
	err := g.service.DeleteAuthor(id)
	if err != nil {
		return nil, fmt.Errorf("could not delete author: %v", err)
	}
	return &pb.DeleteAuthorResponse{Success: true}, nil
}

func toProtoAuthor(author data.Author) *pb.Author {
	return &pb.Author{Id: author.ID, Name: author.Name, PicUrl: author.PicURL}
}

func toAuthorDB(author *pb.Author) data.Author {
	return data.Author{ID: author.Id, Name: author.Name, PicURL: author.PicUrl}
}

// NewGrpcRouter creates a new gRPC router.
func NewGrpcRouter(s Service) GrpcRouter {
	return GrpcRouter{service: s}
}

// CreateAuthor registers a new Author in the database.
func (g GrpcRouter) CreateAuthor(_ context.Context, req *pb.CreateAuthorRequest) (*pb.Author, error) {
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
