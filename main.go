package main

import (
	"github.com/charmbracelet/log"
	"github.com/wcodesoft/mosha-author-service/repository"
	"github.com/wcodesoft/mosha-author-service/service"
	"os"
	"sync"
)

const (
	defaultHttpPort   = "8180"
	defaultGrpcPort   = "8181"
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
	httpPort := getEnv("COMPONENT_PORT", defaultHttpPort)
	mongoHost := getEnv("MONGO_DB_HOST", defaultMongoHost)
	grpcPort := getEnv("GRPC_PORT", defaultGrpcPort)

	mongoClient, err := repository.NewMongoClient(mongoHost)
	if err != nil {
		log.Fatal(err)
	}
	database := repository.NewMongoDatabase(mongoClient, defaultDatabase)
	repo := repository.New(database)
	s := service.New(repo)

	wg := new(sync.WaitGroup)

	wg.Add(2)

	go func() {
		// Create a new HttpRouter.
		router := service.NewHttpRouter(s, AuthorServiceName)
		if err := router.Start(httpPort); err != nil {
			log.Fatal(err)
		}
		wg.Done()
	}()

	go func() {
		grpcRouter := service.NewGrpcRouter(s, AuthorServiceName)
		if err := grpcRouter.Start(grpcPort); err != nil {
			log.Fatal(err)
		}
		wg.Done()
	}()

	wg.Wait()
}
