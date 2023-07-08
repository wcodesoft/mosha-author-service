package repository

import (
	"context"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/wcodesoft/mosha-author-service/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDatabase struct {
	client *mongo.Client
	coll   *mongo.Collection
}

// AddAuthor adds an author to the mongo database.
func (m mongoDatabase) AddAuthor(author data.Author) (string, error) {
	result, err := m.coll.InsertOne(context.Background(), fromAuthor(author))
	if err != nil {
		return "", err
	}
	newId := result.InsertedID
	return fmt.Sprintf("%v", newId), nil
}

// ListAll returns all authors in the mongo database.
func (m mongoDatabase) ListAll() []data.Author {
	cursor, err := m.coll.Find(context.Background(), bson.D{})
	if err != nil {
		panic(err)
	}
	var results []authorDB
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	authors := make([]data.Author, len(results))
	for index, v := range results {
		authors[index] = toAuthor(v)
	}
	return authors
}

// UpdateAuthor updates an author in the mongo database.
func (m mongoDatabase) UpdateAuthor(author data.Author) (data.Author, error) {
	filter := bson.D{{Key: "_id", Value: author.ID}}
	opts := options.Update().SetHint(bson.D{{Key: "_id", Value: 1}})
	update := bson.D{{Key: "$set", Value: fromAuthor(author)}}
	_, err := m.coll.UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		return data.Author{}, err
	}
	return author, nil
}

// DeleteAuthor deletes an author from the mongo database.
func (m mongoDatabase) DeleteAuthor(id string) error {
	filter := bson.D{{Key: "_id", Value: id}}
	opts := options.Delete().SetHint(bson.D{{Key: "_id", Value: 1}})
	result, err := m.coll.DeleteOne(context.Background(), filter, opts)
	if result.DeletedCount == 0 {
		return fmt.Errorf("author with id %s not found", id)
	}
	if err != nil {
		return err
	}
	return nil
}

// GetAuthor returns an author from the mongo database.
func (m mongoDatabase) GetAuthor(id string) (data.Author, error) {
	filter := bson.D{{Key: "_id", Value: id}}
	opts := options.FindOne().SetHint(bson.D{{Key: "_id", Value: 1}})
	var result authorDB
	err := m.coll.FindOne(context.Background(), filter, opts).Decode(&result)
	if err != nil {
		return data.Author{}, err
	}
	return toAuthor(result), nil
}

func NewMongoClient(mongoURI string) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}
	log.Infof("Connected to MongoDB: %s", mongoURI)
	return client, nil
}

func NewMongoDatabase(client *mongo.Client, database string) Database {
	coll := client.Database(database).Collection("authors")
	return &mongoDatabase{
		client: client,
		coll:   coll,
	}
}
