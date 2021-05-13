package main

import (
	"fmt"
	"net/http"
	"time"
)

/**
指定 channel 方向
*/
func executeChannel() {
	ch := make(chan string, 1)
	sendS(ch, "Hello World!")
	readS(ch)
}

func sendS(ch chan<- string, message string) {
	fmt.Printf("Sending: %#v\n", message)
	ch <- message
}

func readS(ch <-chan string) {
	fmt.Printf("Receiving: %#v\n", <-ch)
}

func channelF() {
	//channel1()
	start := time.Now()
	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}
	ch := make(chan string, 20)
	for _, api := range apis {
		go CheckApi(api, ch)
	}
	for i := 0; i < len(apis); i++ {
		fmt.Print(<-ch)
	}
	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}

func checkAPI(api string, ch chan string) {
	_, err := http.Get(api)
	if err != nil {
		ch <- fmt.Sprintf("ERROR: %s is down!\n", api)
		return
	}

	ch <- fmt.Sprintf("SUCCESS: %s is up and running!\n", api)
}

func channel1() {
	size := 2
	ch := make(chan string, size)
	send(ch, "one")
	send(ch, "two")
	go send(ch, "three")
	go send(ch, "four")
	fmt.Println("All data sent to the channel ...")
	for i := 0; i < 4; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("Done!")
}

func send(ch chan string, s string) {
	ch <- s
}

// 多路复用

func muti() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go process(ch1)
	go replicate(ch2)
	for i := 0; i < 2; i++ {
		select {
		case process := <-ch1:
			fmt.Println(process)
		case replicate := <-ch2:
			fmt.Println(replicate)
		}
	}
}

func replicate(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Done replicating!"
}

func process(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "Done Processing!"
}
