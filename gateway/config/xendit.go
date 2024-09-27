package config

import (
	"github.com/xendit/xendit-go/v6"
)

var XenditClient *xendit.APIClient

func InitXendit() {
	apiKey := "xnd_development_FTtBOUKRrARxajVxRPQJIQMzw2WcHIN5FcmT9SoTVbdC3R5aFaMwruoXr5btBJu3"
	XenditClient = xendit.NewClient(apiKey)
}
