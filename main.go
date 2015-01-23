package main

import (
	"log"
	"bytes"
	"github.com/gin-gonic/gin"
)

type slackHookMessage struct {
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

/*
func Serve(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")

	log.Println(text)
}


func main() {
	http.HandleFunc("/inbound", Serve)
	http.ListenAndServe(":8888", nil)
}
*/



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
	//server.Run(":"+os.Getenv("PORT"))
	server.Run(":8888")
}

func PostDataSTDOut(g *gin.Context){
	bufbody := new(bytes.Buffer)
	bufbody.ReadFrom(g.Request.Body)
	body := bufbody.String()

	var message slackHookMessage

	//g.BindWith(&message,binding.Form)
	//g.Bind(&message)
	log.Println(message.Text)
	log.Println(message)

	log.Println(g.Request.Header)

	log.Println(body)
}

