package service

import (
	"authorservice/data"
	"context"
	"github.com/go-kit/kit/endpoint"
)

// addAuthorRequest represents the request to add an author.
type addAuthorRequest struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	PicURL string `json:"picUrl"`
}

// getAuthorRequest represents the request to get an author.
type getAuthorRequest struct {
	ID string `json:"id"`
}

// getAuthorResponse represents the response from getting an author.
type getAuthorResponse struct {
	Author data.Author `json:"author"`
}

// addAuthorResponse represents the response from adding an author.
type addAuthorResponse struct {
	ID string `json:"id"`
}

// deleteAuthorRequest represents the request to delete an author.
type deleteAuthorRequest struct {
	ID string `json:"id"`
}

// deleteAuthorResponse represents the response from deleting an author.
type deleteAuthorResponse struct {
	Err error `json:"err,omitempty"`
}

// authorExistRequest represents the request to check if an author exists.
type authorExistRequest struct {
	ID string `json:"id"`
}

// authorExistResponse represents the response from checking if an author exists.
type authorExistResponse struct {
	Exist bool `json:"exist"`
}

// updateAuthorRequest represents the request to update an author.
type updateAuthorRequest struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	PicURL string `json:"picUrl"`
}

// updateAuthorResponse represents the response from updating an author.
type updateAuthorResponse struct {
	Author data.Author `json:"author"`
}

func makeAddAuthorEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addAuthorRequest)
		id, err := s.CreateAuthor(data.Author{
			ID: req.ID, Name: req.Name, PicURL: req.PicURL,
		})

		if err != nil {
			return addAuthorResponse{}, err
		}

		return addAuthorResponse{ID: id}, nil
	}
}

func makeGetAuthorEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getAuthorRequest)
		id := req.ID
		author, err := s.GetAuthor(id)

		if err != nil {
			return data.Author{}, err
		}

		return getAuthorResponse{author}, nil
	}
}

func makeDeleteAuthorEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteAuthorRequest)
		id := req.ID
		err := s.DeleteAuthor(id)

		if err != nil {
			return deleteAuthorResponse{Err: err}, err
		}

		return deleteAuthorResponse{}, nil
	}
}

func makeAuthorExistEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(authorExistRequest)
		id := req.ID
		exist := s.AuthorExist(id)

		return authorExistResponse{Exist: exist}, nil
	}
}

func makeListAllEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.ListAll(), nil
	}
}

func makeUpdateAuthorEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateAuthorRequest)
		id := req.ID
		name := req.Name
		picUrl := req.PicURL
		author, err := s.UpdateAuthor(data.Author{
			ID: id, Name: name, PicURL: picUrl,
		})

		if err != nil {
			return updateAuthorResponse{}, err
		}

		return updateAuthorResponse{Author: author}, nil
	}
}
