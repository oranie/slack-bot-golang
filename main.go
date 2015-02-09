package main

import (

	"log"
	"net/http"
	"os"
	"strings"
	"github.com/oranie/slack-bot-golang/lib"
)

type SlackHookMesage struct {
	Token string  `json:"token"`
	TeamId string  `json:"team_id"`
	ChannelId string  `json:"channel_id"`
	ChannelName string  `json:"channel_name"`
	Timestamp string  `json:"timestamp"`
	UserId  string  `json:"user_id"`
	UserName string  `json:"user_name"`
	Text string  `json:"text"`
	TriggerWord string  `json:"trigger_word"`
}


func BindSlackData(w http.ResponseWriter, r *http.Request)  {
	var postData SlackHookMesage
	postData.Token = r.FormValue("token")
	postData.TeamId = r.FormValue("team_id")
	postData.ChannelId = r.FormValue("channel_id")
	postData.ChannelName = r.FormValue("channel_name")
	postData.Timestamp = r.FormValue("timestamp")
	postData.UserId = r.FormValue("user_id")
	postData.UserName = r.FormValue("user_name")
	postData.Text = r.FormValue("text")
	postData.TriggerWord = r.FormValue("trigger_word")

	log.Println(postData)
	
	query := strings.Fields(postData.Text)
	log.Println("query string:",query[0]," text string",query[1])
	if query[0] == "image" {
		imageUrl,err := lib.FetchImageUrl(query[1])
		log.Println("image url:",imageUrl)
		if err != nil {
			log.Println("not image query")
		}
		cfg,err := lib.ReadConfig()
		lib.SlackPost(imageUrl,cfg)
	}
}

func main() {
	http.HandleFunc("/v1/slack/inbound", BindSlackData)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	//http.ListenAndServe(":8000", nil)
}
