package repository

import (
	"authorservice/data"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDatabase struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func (m mongoDatabase) AddAuthor(author data.Author) (string, error) {
	result, err := m.coll.InsertOne(context.Background(), fromAuthor(author))
	if err != nil {
		return "", err
	}
	newId := result.InsertedID
	return fmt.Sprintf("%v", newId), nil
}

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

func (m mongoDatabase) UpdateAuthor(author data.Author) (data.Author, error) {
	filter := bson.D{{"_id", author.ID}}
	opts := options.Update().SetHint(bson.D{{"_id", 1}})
	update := bson.D{{"$set", fromAuthor(author)}}
	_, err := m.coll.UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		return data.Author{}, err
	}
	return author, nil
}

func (m mongoDatabase) DeleteAuthor(id string) error {
	filter := bson.D{{"_id", id}}
	opts := options.Delete().SetHint(bson.D{{"_id", 1}})
	_, err := m.coll.DeleteOne(context.Background(), filter, opts)
	if err != nil {
		return err
	}
	return nil
}

func (m mongoDatabase) GetAuthor(id string) (data.Author, error) {
	filter := bson.D{{"_id", id}}
	opts := options.FindOne().SetHint(bson.D{{"_id", 1}})
	var result authorDB
	err := m.coll.FindOne(context.Background(), filter, opts).Decode(&result)
	if err != nil {
		return data.Author{}, err
	}
	return toAuthor(result), nil
}

func (m mongoDatabase) AuthorExist(id string) bool {
	filter := bson.D{{"_id", id}}
	opts := options.FindOne().SetHint(bson.D{{"_id", 1}})
	err := m.coll.FindOne(context.Background(), filter, opts).Err()
	return err == nil
}

func NewMongoDatabase(mongoURI string, database string) Database {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	coll := client.Database(database).Collection("authors")
	return &mongoDatabase{
		client: client,
		coll:   coll,
	}
}
