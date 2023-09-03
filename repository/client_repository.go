package repository

import (
	"context"
	"github.com/charmbracelet/log"
	qpb "github.com/wcodesoft/mosha-service-common/protos/quoteservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientRepository interface {
	DeleteAuthorQuotes(authorID string) (bool, error)
}

type ClientsAddress struct {
	QuoteServiceAddress string
}

type clientRepository struct {
	quoteClient qpb.QuoteServiceClient
}

func (c *clientRepository) DeleteAuthorQuotes(authorID string) (bool, error) {
	response, err := c.quoteClient.DeleteAllQuotesByAuthor(context.Background(), &qpb.DeleteQuotesByAuthorRequest{AuthorId: authorID})
	return response.Success, err
}

// NewClientRepository creates a new client repository.
func NewClientRepository(address ClientsAddress) ClientRepository {
	conn, err := grpc.Dial(address.QuoteServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Errorf("Could not connect to AuthorService at: %s", address.QuoteServiceAddress)
		panic(err)
	}
	client := qpb.NewQuoteServiceClient(conn)
	log.Infof("Connected to AuthorService at: %s", address.QuoteServiceAddress)
	return &clientRepository{
		quoteClient: client,
	}
}
