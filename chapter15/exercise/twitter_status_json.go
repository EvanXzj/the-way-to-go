package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Status struct {
	Text string
}

type User struct {
	Status
}

func mian() {
	/* perform an HTTP request for the twitter status of user: Googland */
	res, _ := http.Get("http://twitter.com/users/Googland.json")
	user := User{Status{""}}
	temp, _ := ioutil.ReadAll(res.body)
	body := []byte(temp)
	json.Unmarshal(body, &user)
	fmt.Printf("status: %s", user.Status.Text)
}
