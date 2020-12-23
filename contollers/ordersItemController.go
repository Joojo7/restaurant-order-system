package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/go-playground/validator.v9"
	database "newapi.com/m/database"
	helpers "newapi.com/m/helpers"
	models "newapi.com/m/models"
)

// connect to the database
var v *validator.Validate = validator.New()

//get orderItemCollection
var orderItemCollection *mongo.Collection = database.OpenCollection(database.Client, "orderItem")

//GetOrderItems is the api used to get a multiple orderItems
func GetOrderItems(response http.ResponseWriter, request *http.Request) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	response.Header().Add("Content-Type", "application/json")

	result, err := orderItemCollection.Find(context.TODO(), bson.M{})
	fmt.Print(result)

	defer cancel()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
	}
	var allOrderItems []bson.M
	if err = result.All(ctx, &allOrderItems); err != nil {
		log.Fatal(err)
	}

	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	json.NewEncoder(response).Encode(allOrderItems)

	// response.Write(jsonBytes)
}

//GetOrderItem is the api used to tget a single orderItem
func GetOrderItem(response http.ResponseWriter, request *http.Request) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	response.Header().Add("Content-Type", "application/json")

	params := mux.Vars(request)

	// id, _ := primitive.ObjectIDFromHex(params["id"])

	var orderItem models.OrderItem

	err := orderItemCollection.FindOne(ctx, bson.M{"orderItem_id": params["id"]}).Decode(&orderItem)
	defer cancel()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
	}

	response.Header().Add("Content-Type", "application/json")

	json.NewEncoder(response).Encode(orderItem)

	// response.Write(jsonBytes)
}

//UpdateOrderItem is used to update orderItems
func UpdateOrderItem(response http.ResponseWriter, request *http.Request) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	// check for content type existence and check for json validity
	helpers.ContentTypeValidator(response, request)

	// call MaxRequestValidator to enforce a maximum read of 1MB .
	dec := helpers.MaxRequestValidator(response, request)

	var orderItem models.OrderItem
	err := dec.Decode(&orderItem)
	helpers.PostPatchRequestValidator(response, request, err)

	params := mux.Vars(request)
	filter := bson.M{"orderItem_id": params["id"]}

	var updateObj primitive.D

	if orderItem.Unit_price != nil {
		updateObj = append(updateObj, bson.E{"price", orderItem.Unit_price})
	}

	orderItem.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	upsert := true
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	result, err := orderItemCollection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{"$set", updateObj},
		},
		&opt,
	)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
	}

	defer cancel()

	response.Header().Add("Content-Type", "application/json")
	json.NewEncoder(response).Encode(result)

}

// var validate *validator.Validate

//CreateOrderItem for creating orderItems
func CreateOrderItem(response http.ResponseWriter, request *http.Request) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	//set response format to JSON
	response.Header().Add("Content-Type", "application/json")

	// check for content type existence and check for json validity
	helpers.ContentTypeValidator(response, request)

	// call MaxRequestValidator to enforce a maximum read of 1MB .
	dec := helpers.MaxRequestValidator(response, request)

	var orderItem models.OrderItem
	err1 := dec.Decode(&orderItem)

	//validate body structure
	if !helpers.PostPatchRequestValidator(response, request, err1) {
		return
	}

	//validate existence if request body

	if v.Struct(&orderItem) != nil {
		response.Write([]byte(fmt.Sprintf(v.Struct(&orderItem).Error())))
		return
	}
	//validate body structure

	orderItem.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	orderItem.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	orderItem.ID = primitive.NewObjectID()
	orderItem.Order_item_id = orderItem.ID.Hex()
	var num = toFixed(*orderItem.Unit_price, 2)
	orderItem.Unit_price = &num

	orderItemCollection.InsertOne(ctx, orderItem)
	defer cancel()

	json.NewEncoder(response).Encode(orderItem)

	defer cancel()

}
