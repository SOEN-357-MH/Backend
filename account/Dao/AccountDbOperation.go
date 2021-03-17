package Dao

import (
	"context"
	"fmt"
	"github.com/superDeano/account/Model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

var (
	Db                *mongo.Collection
	AccountCollection = "Account"
	AccountUsername   = "username"
	Fail              = 0
	Success           = 1
)

func AddUser(u *Model.Account) bool {
	if UsernameExist(u.Username) {
		return false
	}

	_, err := Db.InsertOne(context.TODO(), &u)
	if err != nil {
		panic(err.Error())
		return false
	}
	return true
}

func DeleteUser(u string) string {
	t, err := Db.DeleteOne(context.TODO(), bson.M{AccountUsername: u})
	if err != nil {
		return err.Error()
	}
	switch num := int(t.DeletedCount); {
	case num < Success:
		return "Deleted nothing"
	case num > Success:
		return "Deleted more than an account"
	default:
		return "Deleted an Account"
	}
}

func TestDatabaseConnection() int {
	err := Db.Database().Client().Ping(context.TODO(), nil)
	if err != nil {
		return Fail
	}
	return Success
}

func AuthenticateUser(user *Model.Account) bool {
	userSavedInfo := &Model.Account{}
	project := bson.M{"password": 1, "username": 1}
	_ = Db.FindOne(context.TODO(), bson.D{{AccountUsername, user.Username}}, options.FindOne().SetProjection(project)).Decode(&userSavedInfo)

	if userSavedInfo.Username == user.Username {
		return userSavedInfo.Password == user.Password
	}
	return false
}

func UsernameExist(username string) bool {

	projection := bsonx.Doc{{"username", bsonx.Int32(1)}}
	result := &Model.Account{}
	err := Db.FindOne(context.TODO(), bson.M{AccountUsername: username}, options.FindOne().SetProjection(projection)).Decode(&result)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return result.Username != "" && result != nil

}

func GetUserWithUsername(username string) Model.Account {
	user := &Model.Account{}
	projection := bson.M{"username": 1, "email": 1, "firstname": 1, "lastname": 1}
	err := Db.FindOne(context.TODO(), bson.M{AccountUsername: username}, options.FindOne().SetProjection(projection)).Decode(&user)
	if err != nil {
		fmt.Println(err.Error())
	}

	return *user
}
