package tranapi

import (
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
//	"encoding/json"
//	"bytes"
	"github.com/Akim-Delli/feedrover/config"
)


func Post( ramsContentUrl, site string) {

	resp, err := http.PostForm(config.TRAN_URL_MIGRATE,
		url.Values{"origin_url": {ramsContentUrl}, "dest_site": {site}})
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s\n", string(body))
}

