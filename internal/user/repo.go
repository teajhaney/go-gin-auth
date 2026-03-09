package user

import (
	"context"

	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
)



type Repo struct{
	collection *mongo.Collection
}

func NewRepo(db *mongo.Database) *Repo {
	return &Repo{
		collection: db.Collection("users"),
	}

}

func (r *Repo) FindByEmail(ctx context.Context, email string) (User, error) {
	email = strings.ToLower(strings.TrimSpace(email))
	filter := bson.M{"email": email}
	var user User
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return User{}, fmt.Errorf("user not found")
		}
		return User{}, fmt.Errorf("find by email error: %v", err)
	}
	return user, nil
}


func (r *Repo) Create(ctx context.Context, user User) (User,error) {
	res, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return User{}, fmt.Errorf("error creating user: %v", err)
	}
	id,ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return User{}, fmt.Errorf("error creating user: invalid ID type")
	}
	user.ID = id
	return user, nil

}

