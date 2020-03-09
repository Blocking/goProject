package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type MyUser struct {
	ID       int64    `json:"id"`
	Name     string   `json:"name"`
	LastSeen JsonTime `json:"lastSeen"`
}

func main() {
	user := MyUser{1, "Ken", JsonTime{time.Now()}}
	_ = json.NewEncoder(os.Stdout).Encode(
		user,
	)
	bytes, e := json.Marshal(user)
	if e != nil {
		panic(e)
	}
	fmt.Println(string(bytes))

	var jsonBlob = []byte(` {"id":1,"name":"Ken","lastSeen":"2019-11-11 11:23:12"} `)
	myUser := &MyUser{}
	json.Unmarshal(jsonBlob, myUser)
	fmt.Println(myUser)

}

type JsonTime struct {
	time.Time
}

// 实现它的json序列化方法
func (t1 JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t1.Time).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

func (t *JsonTime) UnmarshalJSON(data []byte) error {
	var err error

	tt, err := time.Parse(`"2006-01-02 15:04:05"`, string(data))

	*t = JsonTime{tt}
	if err != nil {
		return err
	}

	return nil
}
