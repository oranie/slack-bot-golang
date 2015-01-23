package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)


func main() {
	server := gin.New()
	server.Use(gin.Recovery())

	api := server.Group("/v1")
	{
		slack := api.Group("/slack")
		{
			slack.POST("/inbound",PostDataSTDOut)
		}
	}
	server.Run(":"+os.Getenv("PORT"))
}

type slackHookMesage struct {
	Token string  `json:"token"`
	TeamId string  `json:"team_id"`
	ChannelId string  `json:"channel_id"`
	ChannelName string  `json:"channel_name"`
	Timestamp float32  `json:"timestamp"`
	UserId  string  `json:"user_id"`
	UserName string  `json:"user_name"`
	Text string  `json:"text"`
	TriggerWord string  `json:"trigger_word"`
}


func PostDataSTDOut(g *gin.Context){
	//var message slackHookMesage
	var message interface {}
	g.Bind(&message)

	log.Println(message)
}
