package service

import (
	"encoding/json"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/wcodesoft/mosha-author-service/data"
	"net/http"
)

type idResponse struct {
	ID string `json:"id"`
}

type HttpRouter struct {
	service     Service
	serviceName string
}

func NewHttpRouter(s Service, serviceName string) HttpRouter {
	return HttpRouter{
		service:     s,
		serviceName: serviceName,
	}
}

func (h HttpRouter) Start(port string) {
	log.Infof("Starting %s http on %s", h.serviceName, port)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), h.MakeHandler()); err != nil {
		log.Fatalf("Unable to start service %q: %s", h.serviceName, err)
	}
}

func (h HttpRouter) MakeHandler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/api/v1/author/all", h.listAllHandler)
	r.Get("/api/v1/author/{id}", h.getAuthorHandler)
	r.Post("/api/v1/author/delete/{id}", h.deleteAuthorHandler)
	r.Post("/api/v1/author/update", h.updateAuthorHandler)
	r.Post("/api/v1/author", h.addAuthorHandler)

	return r
}

func encodeResponse(w http.ResponseWriter, response interface{}) {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.Write([]byte("Error encoding response"))
	}
}

func (h HttpRouter) addAuthorHandler(w http.ResponseWriter, r *http.Request) {
	var request data.Author
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		encodeResponse(w, err)
		return
	}

	resp, err := h.service.CreateAuthor(request)

	if err != nil {
		encodeResponse(w, err)
		return
	}

	encodeResponse(w, idResponse{ID: resp})
}

func (h HttpRouter) getAuthorHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	resp, err := h.service.GetAuthor(id)

	if err != nil {
		encodeResponse(w, err)
		return
	}

	encodeResponse(w, resp)
}

func (h HttpRouter) deleteAuthorHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.service.DeleteAuthor(id)

	if err != nil {
		encodeResponse(w, err)
		return
	}

	encodeResponse(w, idResponse{ID: id})
}

func (h HttpRouter) updateAuthorHandler(w http.ResponseWriter, r *http.Request) {
	var request data.Author
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		encodeResponse(w, err)
		return
	}

	resp, err := h.service.UpdateAuthor(request)

	if err != nil {
		encodeResponse(w, err)
		return
	}

	encodeResponse(w, resp)
}

func (h HttpRouter) listAllHandler(w http.ResponseWriter, _ *http.Request) {
	resp := h.service.ListAll()

	encodeResponse(w, resp)
}
