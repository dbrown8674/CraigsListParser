Simple Craigslist parser

```package main
   func main() {
   	searchResult, _ := CraigsListParser.Search("baltimore", "tla", "makita")
   
   	for _, v := range searchResult.Item {
   		fmt.Println(v.Title, v.Link)
   	}
   }
```