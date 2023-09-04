package main

import (
	"github.com/charmbracelet/log"
	"github.com/wcodesoft/mosha-author-service/repository"
	"github.com/wcodesoft/mosha-author-service/service"
	mdb "github.com/wcodesoft/mosha-service-common/database"
	mgrpc "github.com/wcodesoft/mosha-service-common/grpc"
	mhttp "github.com/wcodesoft/mosha-service-common/http"
	"os"
	"sync"
)

const (
	defaultHttpPort   = "8180"
	defaultGrpcPort   = "8181"
	AuthorServiceName = "AuthorService"
	defaultMongoHost  = "mongodb://localhost:27017"
	defaultDatabase   = "mosha"
	quoteGrpcAddress  = "localhost:8281"
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
	quoteServiceAddress := getEnv("QUOTE_SERVICE_ADDRESS", quoteGrpcAddress)

	quoteGrpcClientInfo := mgrpc.ClientInfo{
		Name:    "QuoteService",
		Address: quoteServiceAddress,
	}
	clientsRepository, err := repository.NewClientRepository(quoteGrpcClientInfo)
	if err != nil {
		log.Fatal(err)
	}

	mongoClient, err := mdb.NewMongoClient(mongoHost)
	if err != nil {
		log.Fatal(err)
	}
	connection := mdb.NewMongoConnection(mongoClient, defaultDatabase, "authors")
	database := repository.NewMongoDatabase(connection)
	repo := repository.New(database, clientsRepository)
	s := service.New(repo)

	wg := new(sync.WaitGroup)

	wg.Add(2)

	go func() {
		// Create a new AuthorService.
		hs := service.AuthorService{
			Service: s,
			Port:    httpPort,
			Name:    AuthorServiceName,
		}
		err := mhttp.StartHttpService(&hs)
		if err != nil {
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
