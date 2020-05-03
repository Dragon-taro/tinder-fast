package functions

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/Dragon-taro/tinder-fast/types"
)

var jsonStr = []byte(`{"token": "", "facebook_id": ""}`)

func HTTPWithBody(path string, token string, method string) ([]byte, error) {
	url := "https://api.gotinder.com/" + path
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal(err)
	}
	body, err := request(req, token)
	time.Sleep(time.Second * 1)
	return body, err
}

// HTTP is a function for sending Http request
func HTTP(path string, token string, method string) ([]byte, error) {
	url := "https://api.gotinder.com/" + path
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	body, err := request(req, token)
	time.Sleep(time.Second * 1)
	return body, err
}

func setHeader(req *http.Request, token string) {
	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("X-Auth-Token", token)
	}
}

func request(req *http.Request, token string) ([]byte, error) {
	setHeader(req, token) // req自体がポインタ
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// Like is a function for sending like
func Like(token string, u types.ResultUser) {
	path := "like/" + string(u.ID)
	HTTP(path, token, "GET")
}
