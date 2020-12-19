package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	database "newapi.com/m/database"
	models "newapi.com/m/models"
)

// connect to the database
var client *mongo.Client = database.DBinstance()

//get collection
var collection *mongo.Collection = client.Database("prototype").Collection("notes")

var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

//-----------------------------------------------------------------------------------------------api to get
func GetNotes(response http.ResponseWriter, request *http.Request) {

	response.Header().Add("Content-Type", "application/json")

	result, err := collection.Find(context.TODO(), bson.M{})

	defer cancel()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
	}
	var allNotes []bson.M
	if err = result.All(ctx, &allNotes); err != nil {
		log.Fatal(err)
	}

	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	json.NewEncoder(response).Encode(allNotes)

	// response.Write(jsonBytes)
}

//-----------------------------------------------------------------------------------------------api to get
func GetNote(response http.ResponseWriter, request *http.Request) {

	response.Header().Add("Content-Type", "application/json")

	params := mux.Vars(request)

	// id, _ := primitive.ObjectIDFromHex(params["id"])

	var note models.Note

	err := collection.FindOne(ctx, bson.M{"note_id": params["id"]}).Decode(&note)
	// defer cancel()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
	}

	response.Header().Add("Content-Type", "application/json")

	json.NewEncoder(response).Encode(note)

	// response.Write(jsonBytes)
}

//-----------------------------------------------------------------------------------------------api to post
func CreateNote(response http.ResponseWriter, request *http.Request) {
	bodyBytes, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
	}

	contentType := request.Header.Get("Content-type")

	if contentType != "application/json" {
		response.WriteHeader(http.StatusUnsupportedMediaType)
		response.Write([]byte(fmt.Sprintf("need content-type 'application/json' but got '%s'", contentType)))
		return
	}

	var note models.Note
	err = json.Unmarshal(bodyBytes, &note)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(err.Error()))
	}

	note.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	note.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	note.ID = primitive.NewObjectID()
	note.Note_id = note.ID.Hex()

	result, _ := collection.InsertOne(ctx, note)
	json.NewEncoder(response).Encode(result)

}

//-----------------------------------------------------------------------------------------------api to post
func UpdateNote(response http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)

	var translatedBody models.Note
	err = json.Unmarshal(body, &translatedBody)
	var updateObj primitive.M

	if translatedBody.Title == "" && translatedBody.Text == "" {
		response.Write([]byte(fmt.Sprintln("Sorry no input inserted")))
		return
	}

	translatedBody.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	if translatedBody.Text != "" {
		updateObj = bson.M{"text": translatedBody.Text, "updated_at": translatedBody.Updated_at}
	}

	if translatedBody.Title != "" {
		updateObj = bson.M{"title": translatedBody.Title, "updated_at": translatedBody.Updated_at}
	}

	params := mux.Vars(request)
	filter := bson.M{"note_id": params["id"]}

	update := bson.M{
		"$set": updateObj,
	}

	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	result := collection.FindOneAndUpdate(ctx, filter, update, &opt)
	if result.Err() != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
	}

	json.NewEncoder(response).Encode(result)

}
