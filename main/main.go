package main

import (
	"CraigsListParser/CraigsListParser"
	"fmt"
)

func main(){
	searchResult, _ := CraigsListParser.Search("baltimore", "tla", "makita")

	for _, v := range searchResult.Item {
		fmt.Println(v.Title, "\t", v.Link)
	}
}
