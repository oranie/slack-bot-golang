package  lib

import (
	"fmt"
	"os"
	rss "github.com/jteeuwen/go-pkg-rss"
)

func GetHatebu() {
	timeout := 5
	keyword := "これはひどい"
	uri := "http://b.hatena.ne.jp/search/tag?safe=off&q=" + keyword + "&mode=rss" + "&users=30&sort=recent"
	feed := rss.New(timeout, true, chanHandler, itemHandler)
	err := feed.Fetch(uri, nil)

	if err != nil {
		fmt.Fprintf(os.Stderr, "[e] %s: %s", uri, err)
		return
	}
}

func chanHandler(feed *rss.Feed, newchannels []*rss.Channel) {
	fmt.Printf("%d new channel(s) in %s\n", len(newchannels), feed.Url)
}

func itemHandler(feed *rss.Feed, ch *rss.Channel, newitems []*rss.Item) {
	fmt.Printf("%d new item(s) in %s\n", len(newitems), feed.Url)
	for _, item := range newitems {
		fmt.Println(item.Title)
		for _, link := range item.Links {
			fmt.Println(link.Href)
		}
		fmt.Println(item.Description)
		fmt.Println("")
	}
}
