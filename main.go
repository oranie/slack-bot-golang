package main

import (

	"log"
	"net/http"
	"os"
)

type slackHookMesage struct {
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
	var postData slackHookMessage
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

}

func main() {
	http.HandleFunc("/v1/slack/inbound", BindSlackData)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
