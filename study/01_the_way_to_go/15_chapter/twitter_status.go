package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type Status struct {
	Text string
}

type User struct {
	XMLName xml.Name
	Status Status
}

func main() {
	response, _ := http.Get("https://twitter.com/users/Googland.xml")

	user := User{xml.Name{"", "user"}, Status{""}}

	bytes := make([]byte, 2048)
	response.Body.Read(bytes)

	xml.Unmarshal(bytes, &user)
	fmt.Printf("status: %s", user.Status.Text)
}
