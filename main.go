package main

import (
	"authorservice/repository"
	"authorservice/service"
	"log"
	"net/http"
	"os"
)

const (
	defaultPort       = "8180"
	AuthorServiceName = "AuthorService"
	defaultMongoHost  = "mongodb://localhost:27017"
	defaultDatabase   = "mosha"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	log.Printf("Starting %s", AuthorServiceName)
	port := getEnv("COMPONENT_PORT", defaultPort)
	mongoHost := getEnv("MONGO_DB_HOST", defaultMongoHost)
	database := repository.NewMongoDatabase(mongoHost, defaultDatabase)
	repo := repository.New(database)
	s := service.New(repo)
	address := ":" + port
	log.Printf("Starting %s on %s", AuthorServiceName, address)
	if err := http.ListenAndServe(address, service.MakeHandler(s)); err != nil {
		log.Fatalf("Unable to start service %q: %s", AuthorServiceName, err)
	}
}
