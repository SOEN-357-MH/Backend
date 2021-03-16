package Dao

import (
	"github.com/go-bongo/bongo"
	"github.com/superDeano/account/Model"
)

var (
	Db                *bongo.Connection
	AccountCollection = "Account"
	Fail = 1
	Success = 0
)

func AddUser(u *Model.Account) {
	err := Db.Collection(AccountCollection).Save(u)
	if err != nil {
		panic(err.Error())
	}
}

func DeleteUser(u string) {
	//err := db.Collection().Delete()
}

func TestDatabaseConnection() int {
	err := Db.Collection(AccountCollection).Connection.Session.Ping()
	if err != nil {
		return Fail
	}
	return Success
}
