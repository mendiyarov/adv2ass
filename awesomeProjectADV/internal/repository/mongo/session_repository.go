package mongo

import (
	"awesomeProjectADV/internal/domain/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SessionRepository interface {
	Create(*model.Session) error
	GetAll() ([]*model.Session, error)
	GetByID(string) (*model.Session, error)
}

type sessionRepository struct {
	collection *mongo.Collection
}

func NewSessionRepository(db *mongo.Database) SessionRepository {
	return &sessionRepository{
		collection: db.Collection("sessions"),
	}
}

func (r *sessionRepository) Create(s *model.Session) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := r.collection.InsertOne(ctx, s)
	return err
}

func (r *sessionRepository) GetAll() ([]*model.Session, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var sessions []*model.Session
	for cursor.Next(ctx) {
		var s model.Session
		if err := cursor.Decode(&s); err != nil {
			return nil, err
		}
		sessions = append(sessions, &s)
	}

	return sessions, nil
}

func (r *sessionRepository) GetByID(id string) (*model.Session, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var s model.Session
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&s)
	return &s, err
}
