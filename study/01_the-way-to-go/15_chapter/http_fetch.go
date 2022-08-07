package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://www.golang.com/")
	checkErrorHan(err)

	data, err := ioutil.ReadAll(res.Body)
	checkErrorHan(err)

	fmt.Printf("Got: %q", string(data))
}

func checkErrorHan(err error) {
	if err != nil {
		log.Fatalf("Get error : %v", err)
	}
}
