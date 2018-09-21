package main

import (
	"encoding/json"
	"log"

	"github.com/Dragon-taro/tinder-fast/functions"
	"github.com/Dragon-taro/tinder-fast/types"
)

func main() {
	body, err := functions.HTTP("auth", "", "POST")
	if err != nil {
		log.Fatal(err)
	}

	var user types.User
	if err := json.Unmarshal(body, &user); err != nil {
		log.Fatal(err)
	}

	ch := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go likeTenUsers(ch, user)
	}
	for i := 1; i <= 10; i++ {
		<-ch
	}
}

func likeTenUsers(ch chan bool, user types.User) {
	body, err := functions.HTTP("user/recs", user.User.APIToken, "GET") // token := user.User.APIToken
	if err != nil {
		log.Fatal(err)
	}

	var users types.Users
	if err := json.Unmarshal(body, &users); err != nil {
		log.Fatal(err)
	}

	c := make(chan string, 10)

	for _, u := range users.Users {
		go functions.Like(c, user.User.APIToken, u)
	}

	for range users.Users {
		log.Print(<-c)
	}

	ch <- true
}
