package main

import (
	"crypto/md5"
	"encoding/hex"
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
	Hash        string
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

	var rssFeed RSS
	err = xml.Unmarshal(body, &rssFeed)
	if err != nil {
		panic(err)
	}

	fmt.Println(rssFeed.Channel.Title)
	fmt.Println(rssFeed.Channel.Description)
	fmt.Println(rssFeed.Channel.Link)

	for _, item := range rssFeed.Channel.Items {
		hash := generateHash(item.Title, item.PubDate)
		item.Hash = hash
		fmt.Println(item.Title)
		//fmt.Println(item.Description)
		//fmt.Println(item.Link)
		fmt.Println(item.PubDate)
		fmt.Println(item.Hash)
		fmt.Println("---------")
	}
}

func generateHash(title, pubDate string) string {
	h := md5.New()
	h.Write([]byte(title + pubDate))
	return hex.EncodeToString(h.Sum(nil))
}
