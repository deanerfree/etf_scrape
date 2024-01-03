package main

// reddit secret 	mGrKPCjc4_reBeNzK3qY1944mVgdwQ
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type Replies struct {
    Kind string `json:"kind"`
    Data struct {
        Children  []struct {
            Kind string `json:"kind"`
            Data struct {
                ID        string `json:"id"`
                Author    string `json:"author"`
                Body      string `json:"body"`
                Permalink string `json:"permalink"`
                Created   float64 `json:"created"`
            } `json:"data"`
        } `json:"children"`
        Before interface{} `json:"before"`
    } `json:"data"`
}

type RedditResponsePost []struct {
    Data struct {
        Children []struct {
            Data struct {
                Subreddit string `json:"subreddit"`
                Selftext  string `json:"selftext"`
                Title     string `json:"title"`
                Created   float64 `json:"created"`
                Permalink string  `json:"permalink"`
                URL       string  `json:"url"`
								Replies interface{} `json:"replies"`
                // Replies   interface {
                    // Kind string `json:"kind"`
                    // Data struct {
                    //     Children  []struct {
                    //         Kind string `json:"kind"`
                    //         Data struct {
                    //             ID        string `json:"id"`
                    //             Author    string `json:"author"`
                    //             Body      string `json:"body"`
                    //             Permalink string `json:"permalink"`
                    //             Created   float64 `json:"created"`
                    //         } `json:"data"`
                    //     } `json:"children"`
                    //     Before interface{} `json:"before"`
                    // } `json:"data"`
                // } `json:"replies"`
                Author         string  `json:"author"`
                CreatedUtc     float64 `json:"created_utc"`
                AuthorFullname string  `json:"author_fullname"`
                Body           string  `json:"body"`
                Name           string  `json:"name"`
            } `json:"data"`
        } `json:"children"`
				Title string `json:"title"`
    } `json:"data"`
}

type RedditResponse struct {
	Data struct {
		Children []struct {
			Data struct {
				Author string `json:"author"`
				Body string `json:"body"`
				Title string `json:"title"`
				URL string `json:"url"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

func main() {
	posts := findLinks()

	if len(posts) > 0 {
		openFoundLinks(posts)
	}
	
}

func openUrl(url string) []byte{
  resp, err := http.Get(url)
  if err != nil {
      panic(err)
  }
  defer resp.Body.Close()
  body, err := io.ReadAll(resp.Body)
  if err != nil {
      panic(err)
  }
  
  return body
}

func findLinks() []map[string]string{
  url := "https://www.reddit.com/r/ETFs/.json"
	posts := []map[string]string{}
  body := openUrl(url)

  var redditResp RedditResponse

  err := json.Unmarshal(body, &redditResp)

  if err != nil {
      panic(err)
  }
		
	id := 0
  for _, child := range redditResp.Data.Children {
		post := map[string]string{"id": strconv.Itoa(id), "title":child.Data.Title,"url": child.Data.URL}
    fmt.Println("ID:", id)
		fmt.Println("Title:", child.Data.Title)
    fmt.Println("URL:", child.Data.URL)
		if strings.Contains(child.Data.URL, "https://i.redd.it/") {
      continue
    }
		posts = append(posts, post)
    fmt.Println(strings.Repeat("-", 50))
		id++
  }
	return posts
}

func openFoundLinks(posts []map[string]string) {
  arrayOfWords := []string{}
	for i, post := range posts {
    fmt.Println("index: ", i)
		fmt.Println("Opening:", post["id"], post["url"]+".json")
		if post["id"] == "5" || i == 5 {
      fmt.Println("break")
			break
		}

    body := openUrl(post["url"]+".json")

		var redditResponsePost RedditResponsePost

		err := json.Unmarshal(body, &redditResponsePost)
		if err != nil {
    	panic(err)
		}
    
    for _, post := range redditResponsePost {
      for _, child := range post.Data.Children {
        fmt.Println("Title:", post.Data.Title)
        fmt.Println("Body:", child.Data.Body)
        fmt.Println("Author:", child.Data.Author)
        // if replies, ok := child.Data.Replies.(string); ok {
        //     // Replies is a string
        //     fmt.Println("Replies:", replies)
        // } else if replies, ok := child.Data.Replies.(map[string]interface{}); ok {
        //     // Replies is a struct
        //     // You can access fields in the struct using the map
        //     fmt.Println("Replies:", replies["data"])
        // }
        capitalizedWords := findCapitalizedWords(child.Data.Body)
        arrayOfWords = append(arrayOfWords, capitalizedWords...)
      }
    }

  fmt.Println("Capitalized words:", arrayOfWords)
  fmt.Println(strings.Repeat("-", 50))
	}
}

func findCapitalizedWords(word string) []string {
    words := strings.Fields(word)
    var capitalizedWords []string
    for _, word := range words {
      exp, _ := regexp.Compile("^[A-Z]{3,5}$")
      if exp.MatchString(word) {
        capitalizedWords = append(capitalizedWords, word)
      }
    }
    return capitalizedWords
}

