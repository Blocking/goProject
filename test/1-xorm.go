package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"time"
)

// https://github.com/go-xorm/xorm
func main() {
	engine, err := xorm.NewEngine("mysql", "root:root_Zxy@tcp(117.78.0.112:3306)/bim")

	if err != nil {
		log.Fatal(err)
	}
	err1 := engine.Sync2(new(Users), new(raw_data))

	if err1 != nil {
		log.Fatal(err1)
	}

	println(engine.SupportInsertMany())

}

type Users struct {
	Id      int64
	Name    string
	Salt    string
	Age     int
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

type raw_data struct {
	Timestamp         int64   `json:"timestamp"`
	CinemaId          string  `json:"cinemaId"`
	CinemaNo          string  `json:"cinemaNo"`
	ReportTime        string  `json:"reportTime"`
	CinemaStatus      int     `json:"cinemaStatus"`
	BusinessDate      string  `json:"businessDate"`
	ScreenCode        string  `json:"screenCode"`
	FilmCode          string  `json:"filmCode"`
	SessionCode       string  `json:"sessionCode"`
	SessionDatetime   string  `json:"sessionDatetime"`
	LocalSalesCount   int     `json:"localSalesCount"`
	LocalRefundCount  int     `json:"localRefundCount"`
	LocalRefund       float64 `json:"localRefund"`
	LocalSales        float64 `json:"localSales"`
	OnlineSalesCount  int     `json:"onlineSalesCount"`
	OnlineRefundCount int     `json:"onlineRefundCount"`
	OnlineRefund      float64 `json:"onlineRefund"`
	OnlineSales       float64 `json:"onlineSales"`
	PastSaleCount     int     `json:"pastSaleCount"`
	PastSales         float64 `json:"pastSales"`
	AdvanceSalesCount int     `json:"advanceSalesCount"`
	AdvanceSales      float64 `json:"advanceSales"`
	LocalService      float64 `json:"localService"`
	Service           float64 `json:"service"`
}
