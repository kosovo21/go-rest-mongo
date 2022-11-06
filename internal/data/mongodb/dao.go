package mongodb

import (
	"context"
	"time"

	"github.com/kosovo21/go-rest-mongo/internal/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDao struct {
	mongoClient    *mongo.Client
	database       string
	userCollection string
}

func New(mongoClient *mongo.Client, database string) *UserDao {
	return &UserDao{
		mongoClient:    mongoClient,
		database:       database,
		userCollection: "user",
	}
}

func (userDao *UserDao) FindUser(username string) (*data.User, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*3)
	defer cancel()

	user := &data.User{}
	err := userDao.mongoClient.Database(userDao.database).Collection(userDao.userCollection).
		FindOne(ctx, bson.D{{Key: "username", Value: username}}).Decode(user)

	if err == mongo.ErrNilDocument {
		return nil, nil
	}

	return user, err
}
