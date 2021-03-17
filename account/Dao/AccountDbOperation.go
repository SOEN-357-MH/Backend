package Dao

import (
	"github.com/go-bongo/bongo"
	"github.com/superDeano/account/Model"
	"gopkg.in/mgo.v2/bson"
)

var (
	Db                *bongo.Connection
	AccountCollection = "Account"
	AccountUsername   = "username"
	Fail              = 0
	Success           = 1
)

func AddUser(u *Model.Account) bool{
	if UsernameExist(u.Username) {
		return false
	}
	err := Db.Collection(AccountCollection).Save(u)
	if err != nil {
		panic(err.Error())
		return false
	}
	return true
}

func DeleteUser(u string) string {
	t, err := Db.Collection(AccountCollection).Delete(bson.M{AccountUsername: u})
	if err != nil {
		return err.Error()
	}
	switch num := t.Removed; {
	case num < Success:
		return "Deleted nothing"
	case num > Success:
		return "Deleted more than an account"
	default:
		return "Deleted an Account"
	}
}

func TestDatabaseConnection() int {
	err := Db.Collection(AccountCollection).Connection.Session.Ping()
	if err != nil {
		return Fail
	}
	return Success
}

func AuthenticateUser(user *Model.Account) bool {
	results := Db.Collection(AccountCollection).Find(bson.M{AccountUsername: user.Username})
	userSavedInfo := &Model.Account{}
	for results.Next(userSavedInfo) {
		if userSavedInfo.Username == user.Username {
			return userSavedInfo.Password == user.Password
		}
	}
	return false
}

func UsernameExist(username string) bool {
	results := Db.Collection(AccountCollection).Find(bson.M{AccountUsername: username})
	result := &Model.Account{}

	for results.Next(result){
		if result.Username == username {
			return true
		}
	}
	return false
}
