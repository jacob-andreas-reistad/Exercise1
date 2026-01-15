// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
)

type request int 

const (
	inc request = iota // starts inc at 0, then dec at 1 and get at 2
	dec
	get
)

func number_server(request_ch chan request, reply_ch chan int) {
	i := 0

	for {
		select {
		case r := <-request_ch:
			switch r {
			case inc:
				i++
			case dec:
				i--
			case get:
				reply_ch <- i
				return
			}
		}
	}
}

func incrementing(reques_ch chan request, done chan bool) {
	for j := 0; j < 1000000; j++ {
		reques_ch <- inc
	}
    done <- true
}

func decrementing(reques_ch chan request, done chan bool) {
	for j := 0; j < 999958; j++ {
		reques_ch <- dec
	}
    done <- true
}

func main() {
	// What does GOMAXPROCS do? What happens if you set it to 1?
	runtime.GOMAXPROCS(1)

	req_ch := make(chan request)
	rep_ch := make(chan int)
    done   := make(chan bool)

	go number_server(req_ch, rep_ch)
	go incrementing(req_ch, done)
	go decrementing(req_ch, done)

    <- done
	<- done
   
	req_ch <- get
	result := <-rep_ch

	Println("The magic number is:", result)
}
