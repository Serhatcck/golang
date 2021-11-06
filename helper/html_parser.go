package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
)

type head struct {
	Content string `xml:",innerxml"`
}

type body struct {
	Content string `xml:",innerxml"`
	P       string `xml:"p"`
}

type html struct {
	Head head `xml:"head"`
	Body body `xml:"body"`
}

func main() {
	b := []byte(`<!DOCTYPE html>
	<html>
		<head>
			<title>
				Title of the document
			</title>
		</head>
		<body>
			body content 
			<p>more content</p>
		</body>
	</html>`)
	h := html{}
	err := xml.NewDecoder(bytes.NewBuffer(b)).Decode(&h)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(h.Body.Content)
	fmt.Println(h.Body.P)
	fmt.Println(h.Head.Content)
}
