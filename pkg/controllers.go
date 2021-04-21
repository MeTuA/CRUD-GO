package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/metua/crud/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollections = Db().Database("customers").Collection("users")

//controller functions
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := userCollections.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(resp.InsertedID)

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	var userInfo primitive.M

	filter := bson.D{{"name", user.Name}}

	var error models.Error

	err = userCollections.FindOne(context.TODO(), filter).Decode(&userInfo)
	if err != nil {
		fmt.Println("There is no " + user.Name)
		error.Description = "There is no " + user.Name + " in DB."
		json.NewEncoder(w).Encode(error)
	}

	json.NewEncoder(w).Encode(userInfo)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.D{{"name", user.Name}}
	update := bson.D{{"$set", bson.D{{"city", user.City}}}}
	after := options.After
	returnAfter := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	var updatedUser primitive.M

	err = userCollections.FindOneAndUpdate(context.TODO(), filter, update, &returnAfter).Decode(&updatedUser)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(updatedUser)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)["id"]

	_id, err := primitive.ObjectIDFromHex(params)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.D{{"_id", _id}}
	opt := options.Delete().SetCollation(&options.Collation{})

	res, err := userCollections.DeleteOne(context.TODO(), filter, opt)

	json.NewEncoder(w).Encode(res.DeletedCount)

}
