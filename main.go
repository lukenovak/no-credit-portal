package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"nocredit/utils"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	utils.InitRouting(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}