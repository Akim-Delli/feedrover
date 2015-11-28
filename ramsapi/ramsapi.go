package ramsapi

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strconv"
)

var WebsiteName string = "www.elle.com"


func GetNumberOfPages() int {
	resp, err := http.Get("http://" + WebsiteName + "/api/json/updated")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var f interface{}
	err = json.Unmarshal(body, &f)
	m := f.(map[string]interface{})

	return int(m["pages"].(float64))

}

func FetchPageNumber(pageNumber int) interface{} {
	resp, err := http.Get("http://" + WebsiteName + "/api/json/updated/page/" + strconv.Itoa(pageNumber))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var f interface{}
	err = json.Unmarshal(body, &f)
	m := f.(map[string]interface{})

	return m["results"]

}

