package service

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/wcodesoft/mosha-author-service/data"
	"net/http"
)

// addAuthorRequest represents the request to add an author.
type addAuthorRequest struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	PicURL string `json:"picUrl"`
}

// getAuthorResponse represents the response from getting an author.
type getAuthorResponse struct {
	Author data.Author `json:"author"`
}

// addAuthorResponse represents the response from adding an author.
type addAuthorResponse struct {
	ID string `json:"id"`
}

// deleteAuthorResponse represents the response from deleting an author.
type deleteAuthorResponse struct {
	Success bool  `json:"success"`
	Err     error `json:"err,omitempty"`
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

type HttpRouter struct {
	service Service
}

func NewHttpRouter(s Service) HttpRouter {
	return HttpRouter{service: s}
}

func (h HttpRouter) MakeHandler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/api/v1/author/all", h.listAll)
	r.Get("/api/v1/author/{id}", h.getAuthorHandler)
	r.Get("/api/v1/author/exist/{id}", h.authorExistHandler)
	r.Post("/api/v1/author/delete/{id}", h.deleteAuthorHandler)
	r.Post("/api/v1/author/update/{id}", h.updateAuthorHandler)
	r.Post("/api/v1/author", h.addAuthorHandler)

	return r
}

func encodeResponse(w http.ResponseWriter, response interface{}) {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.Write([]byte("Error encoding response"))
	}
}

func (h HttpRouter) addAuthorHandler(w http.ResponseWriter, r *http.Request) {
	var request addAuthorRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		encodeResponse(w, err)
		return
	}

	resp, err := h.service.CreateAuthor(data.Author{
		ID: request.ID, Name: request.Name, PicURL: request.PicURL,
	})

	if err != nil {
		encodeResponse(w, err)
		return
	}

	encodeResponse(w, addAuthorResponse{ID: resp})
}

func (h HttpRouter) getAuthorHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	resp, err := h.service.GetAuthor(id)

	if err != nil {
		encodeResponse(w, err)
		return
	}

	encodeResponse(w, getAuthorResponse{Author: resp})
}

func (h HttpRouter) deleteAuthorHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.service.DeleteAuthor(id)

	if err != nil {
		encodeResponse(w, deleteAuthorResponse{Success: false, Err: err})
		return
	}

	encodeResponse(w, deleteAuthorResponse{Success: true})
}

func (h HttpRouter) authorExistHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	resp := h.service.AuthorExist(id)

	encodeResponse(w, authorExistResponse{Exist: resp})
}

func (h HttpRouter) updateAuthorHandler(w http.ResponseWriter, r *http.Request) {
	var request updateAuthorRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		encodeResponse(w, err)
		return
	}

	resp, err := h.service.UpdateAuthor(data.Author{
		ID: request.ID, Name: request.Name, PicURL: request.PicURL,
	})

	if err != nil {
		encodeResponse(w, err)
		return
	}

	encodeResponse(w, updateAuthorResponse{Author: resp})
}

func (h HttpRouter) listAll(w http.ResponseWriter, _ *http.Request) {
	resp := h.service.ListAll()

	encodeResponse(w, resp)
}
