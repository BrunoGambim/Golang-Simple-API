package repositories

import (
	connectionFactory "Simple-API/database/connection"
	model "Simple-API/domain/model"
	"sync"

	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE_NAME     = "LearningMongo_Golang"
	ALBUMS_COLLECTION = "albums"
)

var repositoryInstance *AlbumRepository
var repositoryInstanceError error
var repositoryOnce sync.Once

type AlbumRepository struct {
	sync.Mutex
	collection *mongo.Collection
	context    context.Context
} //comentarios

func NewRepository() (*AlbumRepository, error) {
	repositoryOnce.Do(func() {
		context := context.TODO()
		client, err := connectionFactory.GetMongoClient(context)

		if err != nil {
			repositoryInstance = &AlbumRepository{}
			repositoryInstanceError = err
		}

		repositoryInstance = &AlbumRepository{
			collection: client.Database(DATABASE_NAME).Collection(ALBUMS_COLLECTION),
			context:    context,
		}
		repositoryInstanceError = nil
	})
	return repositoryInstance, repositoryInstanceError
}

func (repository *AlbumRepository) Insert(album model.Album) (string, error) {
	result, err := repository.collection.InsertOne(repository.context, album)
	if err != nil {
		return "", err
	}
	id := result.InsertedID.(primitive.ObjectID).Hex()
	return id, err
}

func (repository *AlbumRepository) UpdateById(album model.Album, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	updater := bson.M{"$set": album}
	_, err = repository.collection.UpdateByID(repository.context, objectId, updater)
	return err
}

func (repository *AlbumRepository) DeleteById(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectId}
	_, err = repository.collection.DeleteOne(repository.context, filter)
	return err
}

func (repository *AlbumRepository) FindById(id string) (model.Album, error) {
	result := model.Album{}

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}

	filter := bson.M{"_id": objectId}
	err = repository.collection.FindOne(repository.context, filter).Decode(&result)
	return result, err
}

func (repository *AlbumRepository) FindAll() ([]model.Album, error) {
	albums := []model.Album{}
	filter := bson.M{}

	cur, err := repository.collection.Find(repository.context, filter)
	if err != nil {
		return albums, err
	}

	for cur.Next(repository.context) {
		album := model.Album{}
		err := cur.Decode(&album)
		if err != nil {
			return albums, err
		}
		albums = append(albums, album)
	}

	return albums, err
}
