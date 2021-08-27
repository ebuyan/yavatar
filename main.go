package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	yavatar := NewYavatar()
	img, err := yavatar.GetImg(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer img.Body.Close()

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Length", fmt.Sprint(img.ContentLength))
	w.Header().Set("Content-Type", img.Header.Get("Content-Type"))
	if _, err = io.Copy(w, img.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{folder}/{key}", handler)
	http.Handle("/", r)
	log.Fatalln(http.ListenAndServe(":8010", nil))
}
