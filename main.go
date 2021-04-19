package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getAvatarPath(r *http.Request) (path string, err error) {
	vars := mux.Vars(r)
	folder, ok := vars["folder"]
	if !ok {
		err = errors.New("Missing required param 'folder'")
		return
	}
	key, ok := vars["key"]
	if !ok {
		err = errors.New("Missing required param 'key'")
		return
	}
	path = folder + "/" + key + "/600x600"
	return
}

func handler(w http.ResponseWriter, r *http.Request) {
	path, err := getAvatarPath(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req, _ := http.NewRequest(http.MethodGet, "https://avatars.yandex.net/get-music-content/"+path, nil)
	client := http.Client{}
	reqImg, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer reqImg.Body.Close()

	w.Header().Set("Content-Length", fmt.Sprint(reqImg.ContentLength))
	w.Header().Set("Content-Type", reqImg.Header.Get("Content-Type"))
	if _, err = io.Copy(w, reqImg.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/avatar/{folder}/{key}", handler)
	http.Handle("/", r)
	log.Fatalln(http.ListenAndServe(":8010", nil))
}
