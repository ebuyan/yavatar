package main

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type Yavatar struct {
	baseUrl    string
	resolution string
	client     http.Client
}

func NewYavatar() Yavatar {
	return Yavatar{"https://avatars.yandex.net/get-music-content/", "600x600", http.Client{}}
}

func (y Yavatar) GetImg(r *http.Request) (res *http.Response, err error) {
	path, err := y.getPath(r)
	if err != nil {
		return
	}
	req, _ := http.NewRequest(http.MethodGet, y.baseUrl+path, nil)
	client := http.Client{}
	res, err = client.Do(req)
	return
}

func (y Yavatar) getPath(r *http.Request) (path string, err error) {
	vars := mux.Vars(r)
	folder, ok := vars["folder"]
	key, okey := vars["key"]
	if !ok || !okey {
		err = errors.New("Url must be like 'http://<domain>/2411511/6e022a79.a.10223194-2'")
		return
	}
	path = folder + "/" + key + "/" + y.resolution
	return
}
