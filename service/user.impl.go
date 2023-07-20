package service

import (
	"context"

	"github.com/golang-docker/db"
	"github.com/golang-docker/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	userCollection *mongo.Collection = db.GetCollection(db.DB, "users")
	hashManager    AuthService       = NewAuthService()
)

type UserImplement struct{}

func NewUserService() UserService {
	return &UserImplement{}
}

func (u *UserImplement) FindOne(field string, params string) (*model.User, error) {

	user := &model.User{}
	err := userCollection.FindOne(context.TODO(), bson.M{field: params}).Decode(user)
	return user, err
}

func (u *UserImplement) Create(payload *model.User) (*model.User, error) {

	passwordHash, _ := hashManager.HashedPassword(payload.Password)
	payload.Id = primitive.NewObjectID()
	payload.Password = passwordHash

	_, err := userCollection.InsertOne(context.Background(), payload)
	return payload, err
}
