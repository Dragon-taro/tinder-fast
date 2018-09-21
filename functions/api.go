package functions

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Dragon-taro/tinder-fast/types"
)

var jsonStr = []byte(`{"facebook_token": "YOUR_FACEBOOK_ACCESS_TOKEN", "facebook_id": "YOUR_FACEBOOK_ID"}`)

// HTTP is a function for sending Http request
func HTTP(path string, token string, method string) ([]byte, error) {
	url := "https://api.gotinder.com/" + path
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal(err)
	}
	body, err := request(req, token)
	return body, err
}

func setHeader(req *http.Request, token string) {
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("User-Agent", "Tinder/3.0.4 (iPhone; iOS 7.1; Scale/2.00)")
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
func Like(c chan string, token string, u types.ResultUser) {
	path := "like/" + string(u.ID)
	_, err := HTTP(path, token, "GET")

	if err == nil {
		c <- u.Name + "さんをLikeしました！"
	}
}
