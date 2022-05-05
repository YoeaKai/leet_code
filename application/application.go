package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/design_twitter"
)

func run() {
	t := design_twitter.Constructor()
	t.PostTweet(1, 5)
	t.PostTweet(1, 3)
	t.PostTweet(1, 101)
	t.PostTweet(1, 13)
	t.PostTweet(1, 10)
	t.PostTweet(1, 2)
	t.PostTweet(1, 94)
	t.PostTweet(1, 505)
	t.PostTweet(1, 333)
	t.PostTweet(1, 22)
	t.PostTweet(1, 11)
	t.GetNewsFeed(1)
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
