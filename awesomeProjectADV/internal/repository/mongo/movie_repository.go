package mongo

import (
	"awesomeProjectADV/internal/domain/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MovieRepository interface {
	Create(movie *model.Movie) error
	GetAll() ([]*model.Movie, error)
	GetByID(id string) (*model.Movie, error)
	Delete(string) error
}

type movieRepository struct {
	collection *mongo.Collection
}

func NewMovieRepository(db *mongo.Database) MovieRepository {
	return &movieRepository{
		collection: db.Collection("movies"),
	}
}

func (r *movieRepository) Create(movie *model.Movie) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, movie)
	return err
}

func (r *movieRepository) GetAll() ([]*model.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var movies []*model.Movie
	for cursor.Next(ctx) {
		var m model.Movie
		if err := cursor.Decode(&m); err != nil {
			return nil, err
		}
		movies = append(movies, &m)
	}

	return movies, nil
}

func (r *movieRepository) GetByID(id string) (*model.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var m model.Movie
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&m)
	return &m, err
}
func (r *movieRepository) Delete(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}
