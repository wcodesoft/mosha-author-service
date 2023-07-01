package service

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	kithttp "github.com/go-kit/kit/transport/http"
	"net/http"
)

func MakeHandler(service Service) http.Handler {

	addAuthorHandler := kithttp.NewServer(
		makeAddAuthorEndpoint(service),
		decodeAddAuthorRequest,
		encodeResponse)

	listAll := kithttp.NewServer(
		makeListAllEndpoint(service),
		decodeVoidRequest,
		encodeResponse)

	getAuthorHandler := kithttp.NewServer(
		makeGetAuthorEndpoint(service),
		decodeGetAuthorRequest,
		encodeResponse,
	)

	deleteAuthorHandler := kithttp.NewServer(
		makeDeleteAuthorEndpoint(service),
		decodeDeleteAuthorRequest,
		encodeResponse,
	)

	authorExistHandler := kithttp.NewServer(
		makeAuthorExistEndpoint(service),
		decodeAuthorExistRequest,
		encodeResponse,
	)

	updateAuthorHandler := kithttp.NewServer(
		makeUpdateAuthorEndpoint(service),
		decodeUpdateAuthorRequest,
		encodeResponse,
	)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/api/v1/author/all", listAll.ServeHTTP)
	r.Get("/api/v1/author/{id}", getAuthorHandler.ServeHTTP)
	r.Get("/api/v1/author/exist/{id}", authorExistHandler.ServeHTTP)
	r.Post("/api/v1/author/delete/{id}", deleteAuthorHandler.ServeHTTP)
	r.Post("/api/v1/author/update/{id}", updateAuthorHandler.ServeHTTP)
	r.Post("/api/v1/author", addAuthorHandler.ServeHTTP)

	return r
}
func decodeAddAuthorRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request addAuthorRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeUpdateAuthorRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request updateAuthorRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeVoidRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return r, nil
}

func decodeGetAuthorRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id := chi.URLParam(r, "id")
	return getAuthorRequest{id}, nil
}

func decodeDeleteAuthorRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id := chi.URLParam(r, "id")
	return deleteAuthorRequest{id}, nil
}

func decodeAuthorExistRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id := chi.URLParam(r, "id")
	return authorExistRequest{id}, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
