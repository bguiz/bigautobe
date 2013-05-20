package main

import (
	"crudapi"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const HOST_NAME = "localhost"
const API_PREFIX = "/v1"
const PORT_NUM = ":8080"

func hello(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello there!"))
}

func main() {
	// storage
	store := crudapi.NewMapStorage()
	store.AddMap("artists")
	store.AddMap("albums")

	api := crudapi.NewDefaultApiMethods(store)

	// router
	r := mux.NewRouter()

	// mounting the API
	crudapi.MountAPI(r.Host(HOST_NAME).PathPrefix(API_PREFIX).Subrouter(), api)

	// custom handler
	r.HandleFunc("/", hello)

	// start listening
	log.Println("server listening on " + HOST_NAME + PORT_NUM)
	log.Println("API on " + HOST_NAME + PORT_NUM + API_PREFIX)

	err := http.ListenAndServe(PORT_NUM, r)
	if err != nil {
		log.Println(err)
	}
}
