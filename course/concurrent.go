package main

import (
	"fmt"
	"net/http"
	"time"
)

func concurrent() {
	ch := make(chan string)

	start := time.Now()
	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}
	for _, api := range apis {
		go CheckApi(api, ch)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Print(<-ch)
	}
	duration := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", duration.Seconds())
}

func CheckApi(api string, ch chan string) {
	_, err := http.Get(api)
	if err != nil {
		ch <- fmt.Sprintf("ERROR: %s is down!\n", api)
		return
	}
	ch <- fmt.Sprintf("SUCCESS: %s is up and running!\n", api)
}
