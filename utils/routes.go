package utils

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func InitRouting(router *mux.Router) {
	router.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir("./no-credit-portal"))))
	router.HandleFunc("/submit/song", addSongRequest)
	router.HandleFunc("/submit/gbp", addGoodBadPuntableItem)
	router.Handle("/", http.FileServer(http.Dir("./no-credit-portal")))

}

// recieves a song request and sends it to be added to a db
func addSongRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	for key, value := range r.Form {
		fmt.Printf("%s = %s\n", key, value[0])
	}
	w.WriteHeader(200)
}

func addGoodBadPuntableItem(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	item := r.Form.Get("gbp_item")
	fmt.Print(item)
}