package main

import "fmt"
import "github.com/PuerkitoBio/goquery"

func scrape(ch chan string, website string) {
	doc, err := goquery.NewDocument(website)

	if err != nil {
		fmt.Println(err)
	}
	length := doc.Find("[style]").Length()
	str := ""
	if length > 0 {
		str = "%s used the style tag %d times, you shoud be ashamed of yourself.\n"
	} else {
		str = "%s used the style tag %d times, congratulations your website is following a standard set 10 years ago.\n"
	}
	ch <- fmt.Sprintf(str, website, length)
}

func main() {
	websites := []string{"http://www.nytimes.com",
		"http://www.walmart.com",
		"http://www.exxonmobil.com",
		"http://www.chevron.com",
		"http://www.berkshirehathaway.com",
		"http://www.google.com",
		"http://www.pranav.io"}

	ch_results := make(chan string)
	for i := 0; i < len(websites); i++ {
		go scrape(ch_results, websites[i])
	}

	for i := 0; i < len(websites); i++ {
		fmt.Printf("%s", <-ch_results)
	}
}
