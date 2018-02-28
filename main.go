// Copyright © 2017 Kazuki Nishikawa
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/kinokoruumu/cmdTwitter/cmd"
	"github.com/kinokoruumu/cmdTwitter/Service"
	"github.com/kinokoruumu/cmdTwitter/Util"
	"fmt"
)

func main() {
	cmd.Execute()

	api := Util.TwitterApiManager()

	//tweets := Service.GetTimeline(api, "user", "10")
	tweets := Service.Search(api, "#ootd", "10")
	//fmt.Printf("%#v", tweets)

	//for _, tweet := range tweets {
	//	fmt.Println(tweet.CreatedAt,"【",tweet.User.Name,"】", " : ", tweet.Text)
	//	fmt.Println("【",tweet.User.Name,"】", " : ", tweet.Text)
	//}

	for _, tweet := range tweets.Statuses {
		fmt.Printf("[%s] %s\n", tweet.CreatedAt, tweet.Text)
	}

	//Service.ReadHandle(api)


	go Service.StreamTweet(api)
}
