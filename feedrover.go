package main

import (
	"fmt"
	"flag"
	"strings"
	"github.com/Sirupsen/logrus"
	"github.com/Akim-Delli/feedrover/ramsapi"
	"github.com/Akim-Delli/feedrover/persistence"
	"github.com/Akim-Delli/feedrover/tranapi"
)

func main() {

	mode := flag.String("mode", "fetch cosmopolitan ", `2 modes : --mode==fetch site | --mode=migrate contentId (eg cosmopolitan.article.1234`)
	flag.Parse()

	fmt.Println("mode:", *mode)

	if *mode == "fetch" {

		site := flag.Arg(0)
		fmt.Println("fetching latest Content url from site: ", site)

		fetchContentUrlFromSite(site)
		fmt.Println("fetching latest Content url from site: ", site)

	} else if *mode == "migrate" {

		contentLongId := flag.Arg(0)
		inputFields := strings.Split(contentLongId, ".")
		fmt.Println("migrating  ", inputFields)
		sendToRover(inputFields)


	} else {
		logrus.Fatal(*mode + " is not a valid mode")
	}
}

func fetchContentUrlFromSite(websiteName string) {
	totalNumberOfPages := ramsapi.GetNumberOfPages(websiteName)
	fmt.Println("Number of Pages:", totalNumberOfPages)

	for i := 0;  i<=totalNumberOfPages; i++ {

		printHeader()

		dataMap := ramsapi.FetchPageNumber(i, websiteName)

		data := dataMap.([]interface{})
		for _, v := range data {
			val := v.(map[string]interface{})
			url := "www." + websiteName + ".com" + "/api/json/" + val["type"].(string) + "." + val["id"].(string)
			content := persistence.Content{
				val["id"].(string),
				val["content_id"].(string),
				val["type"].(string),
				i,
				val["date_unix"].(string),
				url}
			fmt.Printf("|%-12s|%-15s|%-4d (/%-4d)|%-45s|\n", content.Id, content.Content_type, i, totalNumberOfPages, content.Url)
			persistence.Persist(&content)

		}
	}
}

func printHeader() {
	fmt.Println("|------------+---------------+------------+---------------------------------------------|")
	fmt.Println("|      id    |     type      |page (total)|                  url                        |")
	fmt.Println("|------------+---------------+------------+---------------------------------------------|")

}

func sendToRover(inputFields []string) {
	contentUrl := persistence.GetContentUrl(inputFields[0], inputFields[1], inputFields[2])
	fmt.Println("sending To Tran: " + contentUrl)

	fmt.Println("==========================   Tran Response ============================================")
	contentUrl = "http://" + contentUrl
	tranapi.Post(contentUrl, inputFields[0])
	fmt.Println("=======================================================================================")
}



