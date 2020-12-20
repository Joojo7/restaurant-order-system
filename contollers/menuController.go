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
	database "newapi.com/m/database"
	models "newapi.com/m/models"
)

// connect to the database

//get menuCollection
var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

//GetMenus is the api used to get a multiple menus
func GetMenus(response http.ResponseWriter, request *http.Request) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	response.Header().Add("Content-Type", "application/json")

	result, err := menuCollection.Find(context.TODO(), bson.M{})

	defer cancel()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
	}
	var allMenus []bson.M
	if err = result.All(ctx, &allMenus); err != nil {
		log.Fatal(err)
	}

	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	json.NewEncoder(response).Encode(allMenus)

	// response.Write(jsonBytes)
}

//GetMenu is the api used to tget a single menu
func GetMenu(response http.ResponseWriter, request *http.Request) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	response.Header().Add("Content-Type", "application/json")

	params := mux.Vars(request)

	// id, _ := primitive.ObjectIDFromHex(params["id"])

	var menu models.Menu

	err := menuCollection.FindOne(ctx, bson.M{"menu_id": params["id"]}).Decode(&menu)
	defer cancel()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
	}

	response.Header().Add("Content-Type", "application/json")

	json.NewEncoder(response).Encode(menu)

	// response.Write(jsonBytes)
}

//CreateMenu for creating menus
func CreateMenu(response http.ResponseWriter, request *http.Request) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
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
		defer cancel()

		return
	}

	var menu models.Menu
	err = json.Unmarshal(bodyBytes, &menu)
	fmt.Print(bodyBytes)
	fmt.Print(menu)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(err.Error()))
	}

	// startDate, err := time.Parse("2006-01-02 15:04", menu.sta)
	// if err != nil {
	// 	panic(err)
	// }

	// menu.Start_Date = startDate

	var Start_Date, _ = time.Parse(time.RFC3339, menu.Start_Date.Format(time.RFC3339))
	// menu.Start_Date = Start_Date.String()
	menu.Start_Date, _ = time.Parse(time.RFC3339, Start_Date.Format(time.RFC3339))
	var End_Date, _ = time.Parse(time.RFC3339, menu.End_Date.Format(time.RFC3339))
	menu.End_Date = End_Date
	// menu.End_Date = End_Date.Format(time.RFC3339)

	menu.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	menu.ID = primitive.NewObjectID()
	menu.MenuID = menu.ID.Hex()

	result, _ := menuCollection.InsertOne(ctx, menu)
	defer cancel()
	json.NewEncoder(response).Encode(result)

}

//-----------------------------------------------------------------------------------------------api to post
// func UpdateMenu(response http.ResponseWriter, request *http.Request, ctx context.Context, cancel context.CancelFunc) {
// 	body, err := ioutil.ReadAll(request.Body)

// 	var translatedBody models.Menu
// 	err = json.Unmarshal(body, &translatedBody)
// 	var updateObj primitive.M

// 	if translatedBody.Title == "" && translatedBody.Text == "" {
// 		response.Write([]byte(fmt.Sprintln("Sorry no input inserted")))
// 		return
// 	}

// 	translatedBody.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

// 	if translatedBody.Text != "" {
// 		updateObj = bson.M{"text": translatedBody.Text, "updated_at": translatedBody.Updated_at}
// 	}

// 	if translatedBody.Title != "" {
// 		updateObj = bson.M{"title": translatedBody.Title, "updated_at": translatedBody.Updated_at}
// 	}

// 	params := mux.Vars(request)
// 	filter := bson.M{"menu_id": params["id"]}

// 	update := bson.M{
// 		"$set": updateObj,
// 	}

// 	upsert := true
// 	after := options.After
// 	opt := options.FindOneAndUpdateOptions{
// 		ReturnDocument: &after,
// 		Upsert:         &upsert,
// 	}

// 	result := menuCollection.FindOneAndUpdate(ctx, filter, update, &opt)
// 	if result.Err() != nil {
// 		response.WriteHeader(http.StatusInternalServerError)
// 		response.Write([]byte(err.Error()))
// 	}

// 	json.NewEncoder(response).Encode(result)

// }
