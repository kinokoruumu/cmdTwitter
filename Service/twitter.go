package Service

import (
	"github.com/ChimeraCoder/anaconda"
	"log"
	"net/url"
	"fmt"
)

func Tweet(api *anaconda.TwitterApi, text string) {
	_, err := api.PostTweet(text, nil)
	if(err != nil){
		log.Fatal(err)
	}
}

func Search(api *anaconda.TwitterApi, tag string, count string) anaconda.SearchResponse {
	v := url.Values{}
	v.Set("count", count)
	searchResult, _ := api.GetSearch(tag, v)
	return searchResult
}

func GetTimeline(api *anaconda.TwitterApi, from string, count string) []anaconda.Tweet {
	v := url.Values{}
	v.Set("count", count)

	var tweets []anaconda.Tweet
	var err error

	switch from {
	case "user":
		tweets, err = api.GetUserTimeline(v)
		break
	case "home":
		tweets, err = api.GetHomeTimeline(v)
		break
	}
	if err != nil {
		panic(err)
	}
	return tweets
}

func StreamTweet(api *anaconda.TwitterApi) {

	v := url.Values{}

	stream := api.PublicStreamFilter(v)

	for {
		x := <-stream.C
		switch tweet := x.(type) {
		case anaconda.Tweet:
			fmt.Println(tweet.Text)
		default:
			fmt.Println("unkown type(%T): %v \n", x, x)
		}
	}
}