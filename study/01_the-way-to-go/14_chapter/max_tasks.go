package main

import "fmt"

const MAXREQS = 50

var sem = make(chan int, MAXREQS)

type Request2 struct {
	a, b   int
	replyc chan int
}

func process(r *Request2) {
	fmt.Printf("req.a: %v, req.b: %v", r.a, r.b)
}

func handle(r *Request2) {
	sem <- 1
	process(r)
	<-sem
}

func server2(service chan *Request2) {
	for {
		request := <-service
		go handle(request)
	}
}

func main() {
	service := make(chan *Request2)
	go server2(service)
}
