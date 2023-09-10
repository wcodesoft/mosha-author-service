package service

import (
	"encoding/json"
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/wcodesoft/mosha-author-service/data"
	mhttp "github.com/wcodesoft/mosha-service-common/http"

	"net/http"
)

// AuthorService represents the service interface.
type AuthorService struct {
	Service Service
	Name    string
	Port    string
	mhttp.MoshaHttpService
}

// GetName returns the name of the service.
func (as *AuthorService) GetName() string {
	return as.Name
}

// GetPort returns the port of the service.
func (as *AuthorService) GetPort() string {
	return as.Port
}

// MakeHandler creates a handler for the service.
func (as *AuthorService) MakeHandler() http.Handler {
	r := chi.NewRouter()
	sentryHandler := sentryhttp.New(sentryhttp.Options{
		Repanic: true,
	})

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(sentryHandler.Handle)
	r.Get("/api/v1/author/all", as.listAllHandler)
	r.Get("/api/v1/author/{id}", as.createGetAuthorHandler)
	r.Post("/api/v1/author/delete/{id}", as.deleteAuthorHandler)
	r.Post("/api/v1/author/update", as.updateAuthorHandler)
	r.Post("/api/v1/author", as.addAuthorHandler)

	return r
}

func (as *AuthorService) addAuthorHandler(w http.ResponseWriter, r *http.Request) {
	var request data.Author
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		mhttp.EncodeError(w, err)
		return
	}

	resp, err := as.Service.CreateAuthor(request)

	if err != nil {
		mhttp.EncodeError(w, err)
		return
	}

	mhttp.EncodeResponse(w, mhttp.IdResponse{ID: resp})
}

func (as *AuthorService) createGetAuthorHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	resp, err := as.Service.GetAuthor(id)

	if err != nil {
		mhttp.EncodeError(w, err)
		return
	}

	mhttp.EncodeResponse(w, resp)
}

func (as *AuthorService) deleteAuthorHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := as.Service.DeleteAuthor(id)

	if err != nil {
		mhttp.EncodeError(w, err)
		return
	}

	mhttp.EncodeResponse(w, mhttp.IdResponse{ID: id})
}

func (as *AuthorService) updateAuthorHandler(w http.ResponseWriter, r *http.Request) {
	var request data.Author
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		mhttp.EncodeError(w, err)
		return
	}

	resp, err := as.Service.UpdateAuthor(request)

	if err != nil {
		mhttp.EncodeError(w, err)
		return
	}

	mhttp.EncodeResponse(w, resp)
}

func (as *AuthorService) listAllHandler(w http.ResponseWriter, _ *http.Request) {
	resp := as.Service.ListAll()

	mhttp.EncodeResponse(w, resp)
}
