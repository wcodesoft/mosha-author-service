package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/wcodesoft/mosha-author-service/data"
	"github.com/wcodesoft/mosha-author-service/repository"

	faker "github.com/brianvoe/gofakeit/v6"
)

func createHandler() http.Handler {
	memoryDatabase := repository.NewInMemoryDatabase()
	repo := repository.New(memoryDatabase)
	service := New(repo)
	router := NewHttpRouter(service, "AuthorService")
	handler := router.MakeHandler()
	return handler
}

func executeRequest(req *http.Request, handler http.Handler) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr
}

func jsonReaderFactory(in interface{}) io.Reader {
	buf := bytes.NewBuffer(nil)
	_ = json.NewEncoder(buf).Encode(in)
	return buf
}

func TestHttp(t *testing.T) {

	Convey("When adding valid author", t, func() {
		handler := createHandler()
		author := data.NewAuthorBuilder().WithId(faker.UUID()).WithName(faker.Name()).Build()
		req := httptest.NewRequest("POST", "/api/v1/author", jsonReaderFactory(author))
		rr := executeRequest(req, handler)

		Convey("The response should be 200", func() {
			So(rr.Code, ShouldEqual, http.StatusOK)
		})

		Convey("The response should contain the author ID", func() {
			var authorResponse idResponse
			_ = json.NewDecoder(rr.Body).Decode(&authorResponse)
			So(authorResponse.ID, ShouldEqual, author.ID)
		})
	})

	Convey("When adding invalid author", t, func() {
		handler := createHandler()
		author := data.NewAuthorBuilder().WithId(faker.UUID()).WithName(faker.Name()).Build()

		Convey("When author already exist the response should be 500", func() {
			req1 := httptest.NewRequest("POST", "/api/v1/author", jsonReaderFactory(author))
			executeRequest(req1, handler)
			req2 := httptest.NewRequest("POST", "/api/v1/author", jsonReaderFactory(author))
			rr := executeRequest(req2, handler)
			So(rr.Code, ShouldEqual, http.StatusInternalServerError)
		})

		Convey("When author is invalid the response should be 500", func() {
			rr := executeRequest(
				httptest.NewRequest("POST", "/api/v1/author",
					jsonReaderFactory("invalid"),
				),
				handler,
			)
			So(rr.Code, ShouldEqual, http.StatusInternalServerError)
		})
	})

	Convey("When getting author", t, func() {
		handler := createHandler()
		author := data.NewAuthorBuilder().WithId(faker.UUID()).WithName(faker.Name()).Build()
		req := httptest.NewRequest("POST", "/api/v1/author", jsonReaderFactory(author))
		executeRequest(req, handler)

		Convey("When author exist the response should be 200", func() {
			req := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/author/%s", author.ID), nil)
			rr := executeRequest(req, handler)
			So(rr.Code, ShouldEqual, http.StatusOK)
		})

		Convey("When author does not exist the response should be 500", func() {
			req := httptest.NewRequest("GET", "/api/v1/author/456", nil)
			rr := executeRequest(req, handler)
			So(rr.Code, ShouldEqual, http.StatusInternalServerError)
		})
	})

	Convey("When deleting author", t, func() {
		handler := createHandler()
		author := data.NewAuthorBuilder().WithId(faker.UUID()).WithName(faker.Name()).Build()
		req := httptest.NewRequest("POST", "/api/v1/author", jsonReaderFactory(author))
		executeRequest(req, handler)

		Convey("When author exist the response should be 200", func() {
			req := httptest.NewRequest("POST", fmt.Sprintf("/api/v1/author/delete/%s", author.ID), nil)
			rr := executeRequest(req, handler)
			So(rr.Code, ShouldEqual, http.StatusOK)
		})

		Convey("When author does not exist the response should be 500", func() {
			req := httptest.NewRequest("POST", "/api/v1/author/delete/456", nil)
			rr := executeRequest(req, handler)
			So(rr.Code, ShouldEqual, http.StatusInternalServerError)
		})

	})

	Convey("When listing authors", t, func() {
		handler := createHandler()
		author := data.NewAuthorBuilder().WithId(faker.UUID()).WithName(faker.Name()).Build()
		author2 := data.NewAuthorBuilder().WithId(faker.UUID()).WithName(faker.Name()).Build()

		req := httptest.NewRequest("POST", "/api/v1/author", jsonReaderFactory(author))
		req2 := httptest.NewRequest("POST", "/api/v1/author", jsonReaderFactory(author2))

		executeRequest(req, handler)
		executeRequest(req2, handler)

		Convey("When authors exist the response should be 200", func() {
			req := httptest.NewRequest("GET", "/api/v1/author/all", nil)
			rr := executeRequest(req, handler)
			So(rr.Code, ShouldEqual, http.StatusOK)
		})
	})

	Convey("When updating author", t, func() {
		handler := createHandler()
		id := faker.UUID()
		author := data.NewAuthorBuilder().WithId(id).WithName(faker.Name()).Build()
		req := httptest.NewRequest("POST", "/api/v1/author", jsonReaderFactory(author))
		executeRequest(req, handler)

		Convey("When author exist the response should be 200", func() {
			newName := faker.Name()
			author := data.NewAuthorBuilder().WithId(id).WithName(newName).Build()
			req := httptest.NewRequest("POST", "/api/v1/author/update", jsonReaderFactory(author))
			rr := executeRequest(req, handler)
			So(rr.Code, ShouldEqual, http.StatusOK)
			parsedAuthor := data.NewAuthorBuilder().Build()
			json.NewDecoder(rr.Body).Decode(&parsedAuthor)
			So(parsedAuthor.Name, ShouldEqual, newName)
			So(parsedAuthor.ID, ShouldEqual, id)
		})

		Convey("When author does not exist the response should be 500", func() {
			author := data.NewAuthorBuilder().WithId("426").WithName(faker.Name()).Build()
			req := httptest.NewRequest("POST", "/api/v1/author/update", jsonReaderFactory(author))
			rr := executeRequest(req, handler)
			So(rr.Code, ShouldEqual, http.StatusInternalServerError)
		})
	})
}
