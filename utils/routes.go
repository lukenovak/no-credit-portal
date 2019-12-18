package utils

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func InitRouting(router *mux.Router) {
	router.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir("./no-credit-portal"))))
	router.HandleFunc("/submit/song", addSongRequest)
	router.HandleFunc("/submit/gbp", addGoodBadPuntableItem)
	router.Handle("/", http.FileServer(http.Dir("./no-credit-portal")))
	log.Print("Routes have been initialized")
}

// receives a song request and sends it to be added to a db
func addSongRequest(w http.ResponseWriter, r *http.Request) {
	println("got song request")
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	status, err := AddSongToRequestsTable(r.Form.Get("title"), r.Form.Get("artist"))
	if err != nil {
		log.Print("failed to generate a return value while adding song")
	}
	fmt.Printf(string(status))
	w.WriteHeader(200)
}

func addGoodBadPuntableItem(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Print("unable to parse form")
		return
	}
	item := r.Form.Get("gbp_item")
	fmt.Print(item)
}