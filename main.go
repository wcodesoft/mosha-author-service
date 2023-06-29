package main

import (
	"authorservice/repository"
	"authorservice/service"
	"log"
	"net/http"
)

const (
	defaultAddress    = ":8180"
	AuthorServiceName = "AuthorService"
	defaultMongoHost  = "mongodb://localhost:27017"
	defaultDatabase   = "mosha"
)

func main() {
	database := repository.NewMongoDatabase(defaultMongoHost, defaultDatabase)
	repo := repository.New(database)
	s := service.New(repo)
	if err := http.ListenAndServe(defaultAddress, service.MakeHandler(s)); err != nil {
		log.Fatalf("Unable to start service %q: %s", AuthorServiceName, err)
	}
}
