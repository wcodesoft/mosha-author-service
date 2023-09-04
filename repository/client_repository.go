package repository

import (
	"context"
	mgrpc "github.com/wcodesoft/mosha-service-common/grpc"
	qpb "github.com/wcodesoft/mosha-service-common/protos/quoteservice"
)

type ClientRepository interface {
	DeleteAuthorQuotes(authorID string) error
}

type clientRepository struct {
	quoteClient qpb.QuoteServiceClient
}

// DeleteAuthorQuotes deletes all quotes from an author.
func (c *clientRepository) DeleteAuthorQuotes(authorID string) error {
	_, err := c.quoteClient.DeleteAllQuotesByAuthor(context.Background(), &qpb.DeleteQuotesByAuthorRequest{AuthorId: authorID})
	return err
}

// NewClientRepository creates a new client repository.
func NewClientRepository(clientInfo mgrpc.ClientInfo) (ClientRepository, error) {
	conn, err := clientInfo.NewClientConnection()
	if err != nil {
		return nil, err
	}
	client := qpb.NewQuoteServiceClient(conn)
	return &clientRepository{
		quoteClient: client,
	}, nil
}
