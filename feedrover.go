package main

import (
	"fmt";
	"net/http"
	"io/ioutil"
)

func main() {
	resp, err := http.Get("http://hipsterjesus.com/api")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("response Body:", string(body))
}
