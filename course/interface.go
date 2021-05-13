package main

import (
	"fmt"
	"geometry"
	"io"
	"net/http"
	"os"
)

func interface1() {
	var shape geometry.Shape = geometry.Square{Size: 3}
	geometry.PrintInformation(shape)
	shape = geometry.Circle{Radius: 6}
	geometry.PrintInformation(shape)

	rs := geometry.Person{Name: "John Doe", Country: "USA"}
	ab := geometry.Person{Name: "Mark Collins", Country: "United Kingdom"}
	fmt.Printf("%s\n%s\n", rs, ab)

	resp, err := http.Get("https://api.github.com/users/microsoft/repos?page=15&per_page=5")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	//io.Copy(os.Stdout, resp.Body)´

	customWriter := geometry.CustomWriter{}
	io.Copy(customWriter, resp.Body)
	/*db := database{"Go T-Shirt": 25, "Go Jacket": 55}
	log.Fatal(http.ListenAndServe("localhost:8000", db))*/

}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (d database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("URL:", req.URL)
	for item, price := range d {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
