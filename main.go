package main

import (
	"authorservice/repository"
	"authorservice/service"
	"log"
	"net/http"
)

const (
	defaultAddress        = ":8180"
	AuthorServiceName     = "AuthorService"
	AuthorServiceKeyspace = "authorservice"
	defaultScyllaHost     = "localhost:9042"
)

func main() {
	database := repository.NewScyllaDatabase([]string{defaultScyllaHost}, AuthorServiceKeyspace)
	repo := repository.New(database)
	s := service.New(repo)
	if err := http.ListenAndServe(defaultAddress, service.MakeHandler(s)); err != nil {
		log.Fatalf("Unable to start service %q: %s", AuthorServiceName, err)
	}
}
