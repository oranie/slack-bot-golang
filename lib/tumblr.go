package main

import (
	"github.com/mrjones/oauth"
	"log"
	"io/ioutil"
	"github.com/MariaTerzieva/gotumblr"
	"os"
	"github.com/kurrik/oauth1a"
	"encoding/json"
)

type OauthConfig struct {
	AppKey string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	AccessToken string `json:"access_token"`
	SecretToken string  `json:"access_secret"`
	CallBackUrl string  `json:"callback_url"`
}

func init(){
	var oauthConfig OauthConfig
	oauthConfig.AppKey = os.Getenv("app_key")
	oauthConfig.AppSecret = os.Getenv("app_secret")
	oauthConfig.AccessToken = os.Getenv("access_token")
	oauthConfig.SecretToken = os.Getenv("access_secret")
	oauthConfig.CallBackUrl = os.Getenv("callback_url")
	log.Println("oauth config:",oauthConfig)

}

func (o *OauthConfig) TumblrAuth() {
	consumer := oauth.NewConsumer(o.AppKey, o.AppSecret, oauth.ServiceProvider{})

	accessToken := &oauth.AccessToken{}
	response, err := consumer.Get("http://www.tumblr.com/oauth/request_token", nil, accessToken)
	if err != nil{
		log.Println("req error",err)
	}
	log.Println("Response:", response.StatusCode, response.Status,response.Body,accessToken)
	b, err := ioutil.ReadAll(response.Body)
	log.Println(string(b))
}


type TumblrRequest struct {
	service    *oauth1a.Service
	userConfig *oauth1a.UserConfig
	host       string
	apiKey     string
}


func (o *OauthConfig) FetchTumblrInfo() {
	client := gotumblr.NewTumblrRestClient(
		o.AppKey,
		o.AppSecret,
		o.AccessToken,
		o.SecretToken,
		o.CallBackUrl,
		"http://api.tumblr.com")

	var optionParms = map[string]string{
		"limit": "1",
	}

	blogPhotoInfo := client.Posts("mirakui.tumblr.com","photo",optionParms)
	if len(blogPhotoInfo.Posts) != 0 {
		var base_phpto_post gotumblr.PhotoPost
		for i, _ := range blogPhotoInfo.Posts {
			json.Unmarshal(blogPhotoInfo.Posts[i], &base_phpto_post)
			log.Println(base_phpto_post.Photos[0].Alt_sizes[0])
		}
	}

}

func main(){
	var oauthConfig OauthConfig
	oauthConfig.AppKey = os.Getenv("app_key")
	oauthConfig.AppSecret = os.Getenv("app_secret")
	oauthConfig.AccessToken = os.Getenv("access_token")
	oauthConfig.SecretToken = os.Getenv("access_secret")
	oauthConfig.CallBackUrl = os.Getenv("callback_url")
	log.Println("oauth config:",oauthConfig)

	oauthConfig.FetchTumblrInfo()
}
