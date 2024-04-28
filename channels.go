package main

// func main() {
// 	messages := make(chan string, 2)

// 	go func() {
// 		messages <- "ping"
// 		messages <- "ping2"
// 	}()

// 	msg := <-messages
// 	println(msg)

// 	msg2 := <-messages
// 	println(msg2)
// }

// Channel Synchonoization
// func worker(done chan bool) {
// 	println("working...")
// 	time.Sleep(time.Second)
// 	println("done")

// 	done <- true
// }

// func main() {
// 	done := make(chan bool, 1)
// 	go worker(done)

// 	<-done
// }

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)

	println(<-pongs)
}
