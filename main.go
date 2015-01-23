package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"bytes"
	"github.com/gin-gonic/gin/binding"
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
	Token string  `form:"token"`
	TeamId string  `form:"team_id"`
	ChannelId string  `form:"channel_id"`
	ChannelName string  `form:"channel_name"`
	Timestamp float32  `form:"timestamp"`
	UserId  string  `form:"user_id"`
	UserName string  `form:"user_name"`
	Text string  `form:"text"`
	TriggerWord string  `form:"trigger_word"`
}


func PostDataSTDOut(g *gin.Context){
	var message slackHookMesage
	bufbody := new(bytes.Buffer)
	bufbody.ReadFrom(g.Request.Body)
	body := bufbody.String()

	g.BindWith(&message,binding.Form)
	log.Println(message.Text)
	log.Println(message)

	log.Println(body)
}
