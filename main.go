package main

// reddit secret 	mGrKPCjc4_reBeNzK3qY1944mVgdwQ
import (
	"fmt"
  "github.com/deanerfree/etf_scraper/utils"
)

func main() {
  // initial data from reddit
	posts := utils.FindLinks()

	if len(posts) <= 0 {
    fmt.Println("No posts found")
    return
	}
	initialData := utils.OpenFoundLinks(posts)

  // read data from slice
  // count occurences of each word in the array
  fmt.Println("Initial data:", initialData)
  	
}

