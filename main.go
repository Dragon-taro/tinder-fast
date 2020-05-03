package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Dragon-taro/tinder-fast/functions"
	"github.com/Dragon-taro/tinder-fast/types"
)

func main() {
	body, err := functions.HTTPWithBody("v2/auth/login/facebook", "", "POST")
	if err != nil {
		log.Fatal(err)
	}

	var user types.User
	if err := json.Unmarshal(body, &user); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		likeTenUsers(user)
	}
}

func likeTenUsers(user types.User) {
	body, err := functions.HTTP("user/recs", user.User.APIToken, "GET") // token := user.User.APIToken
	if err != nil {
		log.Fatal(err)
	}

	var users types.Users
	if err := json.Unmarshal(body, &users); err != nil {
		log.Fatal(err)
	}

	for _, u := range users.Users {
		functions.Like(user.User.APIToken, u)
		fmt.Println("liked")
	}
}
