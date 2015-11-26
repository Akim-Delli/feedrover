package main

import (
	"fmt"
	"akim/feedrover/ramsapi"
	"akim/feedrover/persistence"
)

func main() {

	fmt.Println("Number of Pages:", ramsapi.GetNumberOfPages())

	fmt.Println("|     id      |     type      |  page  |                          url                   |")
	fmt.Println("|-------------+---------------+--------+------------------------------------------------|")

	dataMap := ramsapi.FetchPageNumber(1)

	data := dataMap.([]interface{})
	for _, v := range data {
		val := v.(map[string]interface{})
		url := ramsapi.WebsiteName + "/api/json/" + val["type"].(string) + "." + val["id"].(string)
		content := persistence.Content{
						   val["id"].(string),
						   val["content_id"].(string),
					       val["type"].(string),
						   "1",
			               val["date_unix"].(string),
			               url}
		fmt.Printf("|%-13s|%-15s|%-8s|%-45s|\n", content.Id, content.Content_type, "1", content.Url)
		persistence.Persist(&content)

	}



}



