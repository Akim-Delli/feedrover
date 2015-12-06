package main

import (
	"fmt"
	"github.com/Akim-Delli/feedrover/ramsapi"
	"github.com/Akim-Delli/feedrover/persistence"
)

func main() {

	totalNumberOfPages := ramsapi.GetNumberOfPages()
	fmt.Println("Number of Pages:", totalNumberOfPages)

	for i := 0;  i<=totalNumberOfPages; i++ {

		fmt.Println("|------------+---------------+------------+----------------------------------------------|")
		fmt.Println("|      id    |     type      |page (total)|                          url                 |")
		fmt.Println("|------------+---------------+------------+----------------------------------------------|")

		dataMap := ramsapi.FetchPageNumber(i)

		data := dataMap.([]interface{})
		for _, v := range data {
			val := v.(map[string]interface{})
			url := ramsapi.WebsiteName + "/api/json/" + val["type"].(string) + "." + val["id"].(string)
			content := persistence.Content{
				val["id"].(string),
				val["content_id"].(string),
				val["type"].(string),
				i,
				val["date_unix"].(string),
				url}
			fmt.Printf("|%-12s|%-15s|%-4d (/%-4d)|%-45s|\n", content.Id, content.Content_type, i,totalNumberOfPages, content.Url)
			persistence.Persist(&content)

		}
	}
}



