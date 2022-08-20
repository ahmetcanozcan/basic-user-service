package repository

import (
	"app/model"
	"context"

	"github.com/ahmetcanozcan/fet"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository interface {
	Insert(ctx context.Context, user *model.User) (string, error)
	UpdateOne(ctx context.Context, user *model.User) error
	FindByID(ctx context.Context, id string) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	FindAll(ctx context.Context) ([]*model.User, error)
	DeleteOne(ctx context.Context, id string) error
}

type userRepository struct {
	coll *mongo.Collection
}

const (
	userCollectionName = "users"
)

func NewUserRepository(db *mongo.Database) (UserRepository, error) {
	coll := db.Collection(userCollectionName)
	ctx := context.Background()

	if _, err := coll.Indexes().CreateOne(ctx, mongo.IndexModel{Keys: bson.M{"id": 1}, Options: options.Index().SetUnique(true)}); err != nil {
		return nil, err
	}

	if _, err := coll.Indexes().CreateOne(ctx, mongo.IndexModel{Keys: bson.M{"email": 1}, Options: options.Index().SetUnique(true)}); err != nil {
		return nil, err
	}

	return &userRepository{coll: coll}, nil
}

func (r *userRepository) Insert(ctx context.Context, user *model.User) (string, error) {
	user.ID = uuid.New().String()
	_, err := r.coll.InsertOne(ctx, user)

	if err != nil {
		return "", err
	}

	return user.ID, nil
}

func (r *userRepository) UpdateOne(ctx context.Context, user *model.User) error {
	filter := fet.Build(fet.Field("id").Is(user.ID))
	return r.update(ctx, filter, user)
}

func (r *userRepository) update(ctx context.Context, filter, update interface{}) error {
	res, err := r.coll.UpdateOne(ctx,
		filter,
		fet.M{
			fet.KeywordSet: update,
		},
	)

	if err != nil {
		return err
	}

	if res.ModifiedCount == 0 {
		return ErrNotModified
	}

	return nil
}

func (r *userRepository) FindByID(ctx context.Context, id string) (*model.User, error) {
	return r.getByFilter(ctx, fet.Field("id").Is(id))
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	return r.getByFilter(ctx, fet.Field("email").Is(email))
}

func (r *userRepository) FindAll(ctx context.Context) ([]*model.User, error) {
	users := make([]*model.User, 0)

	cur, err := r.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) getByFilter(ctx context.Context, filters ...fet.Updater) (*model.User, error) {
	user := new(model.User)

	filter := fet.Build(filters...)

	err := r.coll.FindOne(ctx, filter).Decode(user)

	// ignore no doc errors, handle it with result
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) DeleteOne(ctx context.Context, id string) error {
	filter := fet.Build(fet.Field("id").Is(id))

	res, err := r.coll.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return ErrNotModified
	}

	return nil
}
