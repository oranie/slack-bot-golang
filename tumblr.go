package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/MariaTerzieva/gotumblr"
	"github.com/kurrik/oauth1a"
	"github.com/mrjones/oauth"
	//"strings"
)

type OauthConfig struct {
	AppKey      string `json:"app_key"`
	AppSecret   string `json:"app_secret"`
	AccessToken string `json:"access_token"`
	SecretToken string `json:"access_secret"`
	CallBackUrl string `json:"callback_url"`
}

type ImageInfo struct {
	TotalPost int64
}

func init() {
	var oauthConfig OauthConfig
	oauthConfig.AppKey = os.Getenv("app_key")
	oauthConfig.AppSecret = os.Getenv("app_secret")
	oauthConfig.AccessToken = os.Getenv("access_token")
	oauthConfig.SecretToken = os.Getenv("access_secret")
	oauthConfig.CallBackUrl = os.Getenv("callback_url")
	log.Println("oauth config:", oauthConfig)

}

func (o *OauthConfig) TumblrAuth() {
	consumer := oauth.NewConsumer(o.AppKey, o.AppSecret, oauth.ServiceProvider{})

	accessToken := &oauth.AccessToken{}
	response, err := consumer.Get("http://www.tumblr.com/oauth/request_token", nil, accessToken)
	if err != nil {
		log.Println("req error", err)
	}
	log.Println("Response:", response.StatusCode, response.Status, response.Body, accessToken)
	b, err := ioutil.ReadAll(response.Body)
	log.Println(string(b))
}

type TumblrRequest struct {
	service    *oauth1a.Service
	userConfig *oauth1a.UserConfig
	host       string
	apiKey     string
}

func (o *OauthConfig) FetchTumblrInfo() int64 {
	client := gotumblr.NewTumblrRestClient(
		o.AppKey,
		o.AppSecret,
		o.AccessToken,
		o.SecretToken,
		o.CallBackUrl,
		"http://api.tumblr.com")

	var optionParms = map[string]string{
		"limit":  "1",
		"offset": "10",
	}

	blogPhotoInfo := client.Posts("mirakui.tumblr.com", "photo", optionParms)
	log.Println("Total Posts:", blogPhotoInfo.Total_posts)
	return blogPhotoInfo.Total_posts
}

func (o *OauthConfig) FetchTumblrImageUrl() string {
	client := gotumblr.NewTumblrRestClient(
		o.AppKey,
		o.AppSecret,
		o.AccessToken,
		o.SecretToken,
		o.CallBackUrl,
		"http://api.tumblr.com")

	TotalPosts := o.FetchTumblrInfo()
	rand.Seed(time.Now().Unix())
	randOffset := rand.Int63n(TotalPosts)

	randOffsetStr := fmt.Sprint(randOffset)
	log.Println("rand offset:", randOffsetStr)

	var optionParms = map[string]string{
		"limit":  "1",
		"offset": randOffsetStr,
	}
	log.Println("options params:", optionParms)

	blogPhotoInfo := client.Posts("mirakui.tumblr.com", "photo", optionParms)

	var imageUrl string
	if len(blogPhotoInfo.Posts) != 0 {
		var base_phpto_post gotumblr.PhotoPost
		for i, _ := range blogPhotoInfo.Posts {
			json.Unmarshal(blogPhotoInfo.Posts[i], &base_phpto_post)
			imageUrl = base_phpto_post.Photos[0].Alt_sizes[0].Url
			log.Println(base_phpto_post.Photos[0].Alt_sizes[0].Url)
		}
	}

	return imageUrl
}

func main() {
	var oauthConfig OauthConfig
	oauthConfig.AppKey = os.Getenv("app_key")
	oauthConfig.AppSecret = os.Getenv("app_secret")
	oauthConfig.AccessToken = os.Getenv("access_token")
	oauthConfig.SecretToken = os.Getenv("access_secret")
	oauthConfig.CallBackUrl = os.Getenv("callback_url")
	log.Println("oauth config:", oauthConfig)

	oauthConfig.FetchTumblrImageUrl()
}
