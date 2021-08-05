package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"ncu-main-recruitment/internal"
	"ncu-main-recruitment/internal/model"
	"ncu-main-recruitment/internal/storage"

	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

var sut internal.App
var server *http.Server

func TestMain(m *testing.M) {
	db := storage.NewDB()
	sut = internal.App{
		DB: db,
	}

	sut.Initialize()
	server = &http.Server{Addr: ":9807", Handler: sut.Router}

	code := m.Run()
	os.Exit(code)
}

func TestPostArticles_ShouldStoreArticlesAgainstUserID(t *testing.T) {
	flushDB()

	data := readSeedData()
	payload, _ := json.Marshal(data)

	request := httptest.NewRequest(http.MethodPost, "/articles", bytes.NewBuffer(payload))

	rw := httptest.NewRecorder()
	server.Handler.ServeHTTP(rw, request)

	article, err := sut.DB.Select(data.UserID)

	if err != nil {
		t.Fatalf("article was not inserted by request")
	}

	if rw.Code != http.StatusCreated {
		t.Fatalf("code is not status created")
	}

	if !reflect.DeepEqual(data, article) {
		t.Fatalf("data and article are not equal")
	}
}

func TestPostArticles_ShouldReturnBadRequestWhenArticlesAlreadyStoredAgainstUserID(t *testing.T) {
	flushDB()

	data := readSeedData()
	payload, _ := json.Marshal(data)

	request := httptest.NewRequest(http.MethodPost, "/articles", bytes.NewBuffer(payload))

	rw := httptest.NewRecorder()
	server.Handler.ServeHTTP(rw, request)

	if rw.Code != http.StatusCreated {
		t.Fatalf("data not stored correctly")
	}

	request = httptest.NewRequest(http.MethodPost, "/articles", bytes.NewBuffer(payload))

	rw = httptest.NewRecorder()
	server.Handler.ServeHTTP(rw, request)

	if rw.Code != http.StatusBadRequest {
		t.Fatalf("code is not status bad request")
	}
}

func flushDB() {
	sut.DB = storage.NewDB()
}

func readSeedData() (articles model.Articles) {
	data, _ := ioutil.ReadFile("./data/seed.json")
	json.Unmarshal(data, &articles)
	return
}
