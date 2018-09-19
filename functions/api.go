package functions

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Dragon-taro/tinder-fist/types"
)

var jsonStr = []byte(`{"facebook_token": "YOUR_FACEBOOK_ACCESS_TOKEN", "facebook_id": "YOUR_FACEBOOK_ID"}`)

func Http(path string, token string, method string) ([]byte, error) {
	// note: tokenのnil許容の方法 -> *stringだとreq.Header.Setでエラー
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

	body, err := ioutil.ReadAll(resp.Body)

	return body, err
}

func Like(c chan string, token string, u types.ResultUser) {
	path := "like/" + string(u.ID)
	_, err := Http(path, token, "GET")

	if err == nil {
		c <- u.Name + "さんをLikeしました！"
	}
}
