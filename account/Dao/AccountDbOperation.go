package Dao

import (
	"context"
	"fmt"
	"github.com/superDeano/account/Model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"log"
	"net/http"
)

var (
	Db                *mongo.Collection
	AccountCollection = "Account"
	AccountUsername   = "username"
	AccountEmail      = "email"
	AccountPassword   = "password"
	AccountLastName   = "lastname"
	AccountFirstName  = "firstname"
	ShowId            = "showId"
	MovieId           = "movieId"
	Movie             = 1
	Show              = 0
	movieWatchList    = "movie_watchlist"
	showWatchList     = "show_watchlist"
	Fail              = -1
	Success           = 1
)

func AddUser(u *Model.Account) bool {
	if CheckUsernameOrEmailInUser(u.Username, u.Email) {
		return false
	}
	u.MovieWatchlist = make([]int, 0)
	u.ShowWatchlist = make([]int, 0)
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

func UsernameExistsEh(username string) bool {

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

func AddShowToWatchList(username string, showId int) (int, string) {
	return addMediaToWatchlist(username, Show, showId)
}

func AddMovieToWatchlist(username string, movieId int) (int, string) {
	return addMediaToWatchlist(username, Movie, movieId)
}

func addMediaToWatchlist(username string, mediaType int, id int) (int, string) {
	match := bson.M{AccountUsername: username}
	var insert bson.M
	switch mediaType {
	case Movie:
		insert = bson.M{"$addToSet": bson.M{movieWatchList: id}}
	default:
		insert = bson.M{"$addToSet": bson.M{showWatchList: id}}
	}
	res, err := Db.UpdateOne(context.TODO(), match, insert)

	if err != nil {
		log.Printf(err.Error())
		return http.StatusExpectationFailed, err.Error()
	}
	if res.ModifiedCount != 1 {
		return http.StatusExpectationFailed, "Did not add"
	}
	return http.StatusOK, "Added!"
}

func GetShowWatchlist(username string) (int, []int) {
	if UsernameExistsEh(username) {
		return http.StatusOK, getWatchlist(username, Show)
	} else {
		return http.StatusNotFound, nil
	}
}

func GetMovieWatchlist(username string) (int, []int) {
	if UsernameExistsEh(username) {
		return http.StatusOK, getWatchlist(username, Movie)
	} else {
		return http.StatusNotFound, nil
	}
}

func getWatchlist(username string, mediaType int) []int {
	var projection bson.M
	var watchlist Model.Account
	switch mediaType {
	case Movie:
		projection = bson.M{movieWatchList: 1, "_id": 0}
		err := Db.FindOne(context.TODO(), bson.M{AccountUsername: username}, options.FindOne().SetProjection(projection)).Decode(&watchlist)
		if err != nil {
			log.Println(err.Error())
		}
		return watchlist.MovieWatchlist
	default:
		projection = bson.M{showWatchList: 1, "_id": 0}
		err := Db.FindOne(context.TODO(), bson.M{AccountUsername: username}, options.FindOne().SetProjection(projection)).Decode(&watchlist)
		if err != nil {
			log.Println(err.Error())
		}
		return watchlist.ShowWatchlist
	}

}

func RemoveMovieFromWatchlist(username string, id int) (int, string) {
	if UsernameExistsEh(username) {
		return removeFromWatchlist(username, Movie, id)
	} else {
		return http.StatusNotFound, "User does not exist"
	}
}

func RemoveShowFromWatchlist(username string, id int) (int, string) {
	if UsernameExistsEh(username) {
		return removeFromWatchlist(username, Show, id)
	} else {
		return http.StatusNotFound, "User does not exist"
	}
}

func removeFromWatchlist(username string, mediaType int, id int) (int, string) {
	whoQuery := bson.M{AccountUsername: username}
	var updateQuery bson.M
	switch mediaType {
	case Movie:
		updateQuery = bson.M{"$pull": bson.M{movieWatchList: id}}
	default:
		updateQuery = bson.M{"$pull": bson.M{showWatchList: id}}
	}
	res, err := Db.UpdateOne(context.TODO(), whoQuery, updateQuery)
	if err != nil {
		log.Println(err.Error())
	}
	if res.ModifiedCount != 1 {
		return http.StatusExpectationFailed, "Did not remove only one!"
	} else {
		return http.StatusOK, "Removed!"
	}
}
