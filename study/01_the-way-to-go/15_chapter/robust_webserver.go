package main

import (
	"io"
	"log"
	"net/http"
)

const FORM = `
	<html><body>
		<form action="#" method="post" name="bar">
			<input type="text" name="in">
			<input type="submit" value="Submit">
		</form>
	</body></html>
`

type HandleFnc func(http.ResponseWriter, *http.Request)

func SimpleServer1(w http.ResponseWriter, request *http.Request) {
	panic("test panic")
	io.WriteString(w, "<h1>Hello, world</h1>")
}

func FormServer1(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch request.Method {
	case "GET":
		io.WriteString(w, FORM)
	case "POST":
		io.WriteString(w, request.FormValue("in"))
	}
}

func logPanics(function HandleFnc) HandleFnc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic: %v", request.RemoteAddr, x)
			}
		}()
		function(writer, request)
	}
}

func main() {
	http.HandleFunc("/test1", logPanics(SimpleServer1))
	http.HandleFunc("/test2", logPanics(FormServer1))

	if err := http.ListenAndServe(":8088", nil); err != nil {
		panic(err)
	}
}
