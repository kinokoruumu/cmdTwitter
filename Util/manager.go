package Util

import (
	"github.com/kinokoruumu/cmdTwitter/Constants"
	"github.com/ChimeraCoder/anaconda"
)
func TwitterApiManager() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(Constants.CONSUMER_KEY)
	anaconda.SetConsumerSecret(Constants.CONSUMER_SECRET)
	api := anaconda.NewTwitterApi(Constants.ACCESS_TOKEN, Constants.ACCESS_TOKE_SECRET)
	return api
}