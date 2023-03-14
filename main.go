package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RSS struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
}

func main() {
	url := "https://princeoftravel.com/feed/"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var rss RSS
	err = xml.Unmarshal(body, &rss)
	if err != nil {
		panic(err)
	}

	fmt.Println(rss.Channel.Title)
	fmt.Println(rss.Channel.Description)
	fmt.Println(rss.Channel.Link)

	for _, item := range rss.Channel.Items {
		fmt.Println(item.Title)
		//fmt.Println(item.Description)
		//fmt.Println(item.Link)
		fmt.Println(item.PubDate)
		fmt.Println("---------")
	}
}
