package design_twitter

type Twitter struct {
	userId    []int
	tweetId   []int
	followMap map[int][]int
}

func Constructor() Twitter {
	return Twitter{followMap: make(map[int][]int)}
}

func (twitter *Twitter) PostTweet(userId int, tweetId int) {
	twitter.userId = append(twitter.userId, userId)
	twitter.tweetId = append(twitter.tweetId, tweetId)
}

func (twitter *Twitter) GetNewsFeed(userId int) []int {
	res := make([]int, 0, 10)
	followers := twitter.followMap[userId]
	followers = append(followers, userId)

	for i := len(twitter.tweetId) - 1; i >= 0 && len(res) < 10; i-- {
		var isGet bool

		for j := 0; j < len(followers); j++ {
			if twitter.userId[i] == followers[j] {
				isGet = true
				break
			}
		}

		if isGet {
			res = append(res, twitter.tweetId[i])
		}
	}

	return res
}

func (twitter *Twitter) Follow(followerId int, followeeId int) {
	twitter.followMap[followerId] = append(twitter.followMap[followerId], followeeId)
}

func (twitter *Twitter) Unfollow(followerId int, followeeId int) {
	for i, id := range twitter.followMap[followerId] {
		if id == followeeId {
			twitter.followMap[followerId] = append(twitter.followMap[followerId][:i], twitter.followMap[followerId][i+1:]...)
		}
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
