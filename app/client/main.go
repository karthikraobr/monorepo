package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	req, err := http.Get("localhost:8080/hello")
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer req.Body.Close()
	log.Print(string(b))
	log.Print("just to trigger ci")
}
