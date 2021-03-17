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
	AccountEmail      = "email"
	AccountPassword   = "password"
	AccountLastName   = "lastname"
	AccountFirstName  = "firstname"
	Fail              = -1
	Success           = 1
)

func AddUser(u *Model.Account) bool {
	if CheckUsernameOrEmailInUser(u.Username, u.Email) {
		return false
	}

	_, err := Db.InsertOne(context.TODO(), &u)
	if err != nil {
		panic(err.Error())
		return false
	}
	return true
}

func DeleteUser(u string) int {
	t, err := Db.DeleteMany(context.TODO(), bson.M{AccountUsername: u})
	if err != nil {
		return Fail
	}
	return int(t.DeletedCount)
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
	project := bson.M{AccountPassword: 1, AccountUsername: 1}
	_ = Db.FindOne(context.TODO(), bson.D{{AccountUsername, user.Username}}, options.FindOne().SetProjection(project)).Decode(&userSavedInfo)

	if userSavedInfo.Username == user.Username {
		return userSavedInfo.Password == user.Password
	}
	return false
}

func UsernameExist(username string) bool {

	projection := bsonx.Doc{{AccountUsername, bsonx.Int32(1)}}
	result := &Model.Account{}
	err := Db.FindOne(context.TODO(), bson.M{AccountUsername: username}, options.FindOne().SetProjection(projection)).Decode(&result)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return result.Username != "" && result != nil

}

func GetUserWithUsername(username string) (Model.Account, error) {
	user := &Model.Account{}
	projection := bson.M{AccountUsername: 1, AccountEmail: 1, AccountFirstName: 1, AccountLastName: 1}
	err := Db.FindOne(context.TODO(), bson.M{AccountUsername: username}, options.FindOne().SetProjection(projection)).Decode(&user)
	if err != nil {
		fmt.Println(err.Error())
	}
	return *user, err
}

func CheckUsernameOrEmailInUser(username string, email string) bool {
	user := &Model.Account{}
	projection := bson.M{AccountUsername: 1, AccountEmail: 1}
	findQuery := bson.M{}
	var orQuery []bson.M
	orQuery = append(orQuery, bson.M{AccountEmail: email}, bson.M{AccountUsername: username})

	findQuery["$or"] = orQuery

	err := Db.FindOne(context.TODO(), findQuery, options.FindOne().SetProjection(projection)).Decode(&user)
	if err != nil {
		return false
	}
	return user.Username != "" || user.Email != ""
}
