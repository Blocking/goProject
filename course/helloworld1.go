package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	/*var p = reader("/tmp/dsapi.xuban.v1-data.csv")
	for i,v := range p {
		fmt.Println(i,v[1])
	}
	fmt.Println(p)*/
	var st = "accessToken=0a7a8ef3-9317-4e41-b3f9-528e2122f680&apikey=9QwbZ9RK5ZpcRxM1CPNZeNtm1dTr4Lip&eventApp=pp&eventDeviceType=android&eventVersion=v1.0&quarterCode=8&schoolId=1&studentCode=BJ1120470193&t=1616688011701&userId=224f52e6-4ba3-4ecd-aa7a-b4737994bef7"
	arr := strings.Split(st, "&")
	for _, v := range arr {
		param := strings.Split(v, "=")
		fmt.Println(param)
	}
	fmt.Println(arr)
}
func reader(path string) [][]string {
	csvFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()
	csvReader := csv.NewReader(csvFile)
	all, err1 := csvReader.ReadAll()
	if err1 != nil {
		panic(err1)
	}
	return all
}

func writer(data [][]string) {
	csvFile, err := os.Create("/tmp/bar.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()
	csvWriter := csv.NewWriter(csvFile)
	rows := [][]string{
		{"123", "John", "john@example.com", "$141,987"},
		{"456", "Sam", "sam@example.com", "$905,234"},
		{"678", "Jane", "jane@example.com", "$548,980"},
	}
	err = csvWriter.WriteAll(rows)
	if err != nil {
		fmt.Printf("error (%v)", err)
		return
	}

}
