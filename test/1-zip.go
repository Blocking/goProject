package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"io"
	"log"
	"time"
)

func main() {
	// Open a zip archive for reading.

	r, err := zip.OpenReader("D:\\tmp\\statisticBoxOfice2019-10-31.181000.757.json.zip")

	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	engine := getMysqlConnectEngine()

	/*err1 := engine.Sync2(new(RawData))

	if err1 != nil{
		log.Fatal(err1)
	}*/

	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.File {
		fmt.Printf("Contents of %s\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		bytes := make([]byte, f.FileInfo().Size())
		_, err = io.ReadFull(rc, bytes)
		fmt.Println(string(bytes))

		var data []RawData
		err = json.Unmarshal(bytes, &data)
		if err != nil {
			log.Fatal(err)
		}
		rc.Close()
		fmt.Println(len(data))
		_, err = engine.Insert(data)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func getMysqlConnectEngine() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:root@tcp(localhost:3306)/customers")

	if err != nil {
		log.Fatal(err)
	}
	return engine
}

type RawData struct {
	Timestamp         int64    `json:"timestamp"`
	CinemaId          string   `json:"cinemaId"`
	CinemaNo          string   `json:"cinemaNo"`
	ReportTime        JSONTime `json:"reportTime"`
	CinemaStatus      int      `json:"cinemaStatus"`
	BusinessDate      string   `json:"businessDate"`
	ScreenCode        string   `json:"screenCode"`
	FilmCode          string   `json:"filmCode"`
	SessionCode       string   `json:"sessionCode"`
	SessionDatetime   string   `json:"sessionDatetime"`
	LocalSalesCount   int      `json:"localSalesCount"`
	LocalRefundCount  int      `json:"localRefundCount"`
	LocalRefund       float64  `json:"localRefund"`
	LocalSales        float64  `json:"localSales"`
	OnlineSalesCount  int      `json:"onlineSalesCount"`
	OnlineRefundCount int      `json:"onlineRefundCount"`
	OnlineRefund      float64  `json:"onlineRefund"`
	OnlineSales       float64  `json:"onlineSales"`
	PastSaleCount     int      `json:"pastSaleCount"`
	PastSales         float64  `json:"pastSales"`
	AdvanceSalesCount int      `json:"advanceSalesCount"`
	AdvanceSales      float64  `json:"advanceSales"`
	LocalService      float64  `json:"localService"`
	Service           float64  `json:"service"`
}

type JSONTime struct {
	time.Time
}

// 实现它的json序列化方法
func (t1 JSONTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t1.Time).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

func (t *JSONTime) UnmarshalJSON(data []byte) error {
	var err error

	tt, err := time.Parse(`"2006-01-02 15:04:05"`, string(data))

	*t = JSONTime{tt}
	if err != nil {
		return err
	}

	return nil
}
