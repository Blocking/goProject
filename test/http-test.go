package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://gjdyzjb.cn/bits/w/users/s?page=0&size=50&sort=id,desc", nil)
	if err != nil {
		log.Fatalln(err)
	}

	cookie := http.Cookie{Name: "SESSION", Value: "f4116174-ea7d-40ab-8cb1-5bbdbd148bbc", Expires: time.Now().Add(111 * time.Second)}

	req.AddCookie(&cookie)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
	data := ResponseData{}

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(data.Status)
	log.Println(len(data.Data))
	log.Println(data.Data[0].Name)
	log.Println(data.Data[0].Password)
	log.Println(data)
}

// 依据json 生成对应 struct结构体  生成转换地址：https://mholt.github.io/json-to-go/
type ResponseData struct {
	Status string `json:"status"`
	Data   []struct {
		ID                    int         `json:"id"`
		No                    string      `json:"no"`
		Password              interface{} `json:"password"`
		Name                  string      `json:"name"`
		UserType              string      `json:"userType"`
		OrgID                 interface{} `json:"orgId"`
		ProvinceCode          interface{} `json:"provinceCode"`
		CityCode              interface{} `json:"cityCode"`
		CountyCode            interface{} `json:"countyCode"`
		Mobile                string      `json:"mobile"`
		Email                 interface{} `json:"email"`
		ContactUser           interface{} `json:"contactUser"`
		ResidentID            int         `json:"residentId"`
		Actived               bool        `json:"actived"`
		ForceCertificateLogin bool        `json:"forceCertificateLogin"`
		RoleIds               []int       `json:"roleIds"`
		Roles                 string      `json:"roles"`
		Locked                bool        `json:"locked"`
	} `json:"data"`
	Pageable struct {
		TotalElements    int  `json:"totalElements"`
		NumberOfElements int  `json:"numberOfElements"`
		TotalPages       int  `json:"totalPages"`
		Number           int  `json:"number"`
		First            bool `json:"first"`
		Last             bool `json:"last"`
		Size             int  `json:"size"`
		FromNumber       int  `json:"fromNumber"`
		ToNumber         int  `json:"toNumber"`
	} `json:"pageable"`
}
