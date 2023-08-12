package service

import (
	"encoding/json"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/wcodesoft/mosha-author-service/data"
	mhttp "github.com/wcodesoft/mosha-service-common/http"
	"net/http"
	"time"
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

func (h *HttpRouter) Start(port string) error {
	log.Infof("Starting %s http on %s", h.serviceName, port)

	server := &http.Server{
		Addr:              fmt.Sprintf(":%s", port),
		Handler:           h.MakeHandler(),
		ReadHeaderTimeout: 3 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		return fmt.Errorf("unable to start service %q: %s", h.serviceName, err)
	}
	return nil
}

func (h *HttpRouter) MakeHandler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/api/v1/author/all", h.listAllHandler)
	r.Get("/api/v1/author/{id}", h.createGetAuthorHandler)
	r.Post("/api/v1/author/delete/{id}", h.deleteAuthorHandler)
	r.Post("/api/v1/author/update", h.updateAuthorHandler)
	r.Post("/api/v1/author", h.addAuthorHandler)

	return r
}

func (h *HttpRouter) addAuthorHandler(w http.ResponseWriter, r *http.Request) {
	var request data.Author
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		mhttp.EncodeError(w, err)
		return
	}

	resp, err := h.service.CreateAuthor(request)

	if err != nil {
		mhttp.EncodeError(w, err)
		return
	}

	mhttp.EncodeResponse(w, idResponse{ID: resp})
}

func (h *HttpRouter) createGetAuthorHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	resp, err := h.service.GetAuthor(id)

	if err != nil {
		mhttp.EncodeError(w, err)
		return
	}

	mhttp.EncodeResponse(w, resp)
}

func (h *HttpRouter) deleteAuthorHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.service.DeleteAuthor(id)

	if err != nil {
		mhttp.EncodeError(w, err)
		return
	}

	mhttp.EncodeResponse(w, idResponse{ID: id})
}

func (h *HttpRouter) updateAuthorHandler(w http.ResponseWriter, r *http.Request) {
	var request data.Author
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		mhttp.EncodeError(w, err)
		return
	}

	resp, err := h.service.UpdateAuthor(request)

	if err != nil {
		mhttp.EncodeError(w, err)
		return
	}

	mhttp.EncodeResponse(w, resp)
}

func (h *HttpRouter) listAllHandler(w http.ResponseWriter, _ *http.Request) {
	resp := h.service.ListAll()

	mhttp.EncodeResponse(w, resp)
}
