package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"newapi.com/m/models"
)

type Coaster struct {
	Text         string
	Manufacturer string
	ID           string
	InPark       string
	Heignt       int
}

type CoasterHandlers struct {
	sync.Mutex
	store map[string]Coaster
}

type NoteHandlers struct {
	sync.Mutex
	store map[string]models.Note
}

//-----------------------------------------------------------------------------------------------api to get
func (h *CoasterHandlers) Get(response http.ResponseWriter, request *http.Request) {
	coasters := make([]Coaster, len(h.store))

	h.Lock()
	i := 0
	for _, coaster := range h.store {
		coasters[i] = coaster
		i++
	}
	h.Unlock()

	jsonBytes, err := json.Marshal(coasters)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
	}

	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(jsonBytes)
}

//-----------------------------------------------------------------------------------------------api to get
// func (h *coasterHandlers) getCoaster(response http.ResponseWriter, request *http.Request) {

// 	request.URL
// }

//-----------------------------------------------------------------------------------------------api to post
func (h *CoasterHandlers) Post(response http.ResponseWriter, request *http.Request) {
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

	var coaster Coaster
	err = json.Unmarshal(bodyBytes, &coaster)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(err.Error()))
	}

	coaster.ID = fmt.Sprintf("%d", time.Now().UnixNano())

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, coaster)
	json.NewEncoder(response).Encode(result)

	h.Lock()
	h.store[coaster.ID] = coaster
	response.Write([]byte(fmt.Sprintf("object has been created: '%v'", coaster)))
	defer h.Unlock()
}

//NewCoasterHandlers func
func NewCoasterHandlers() *CoasterHandlers {
	return &CoasterHandlers{
		store: map[string]Coaster{},
	}
}
