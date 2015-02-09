package lib

import (
	"os"
	"net/url"
	"net/http"
	"log"
	"encoding/json"
	"errors"
	"os/user"
)

type SlackMsg struct {
	Channel string `json:"channel"`
	Username string `json:"username,omitempty"`
	Text string `json:"text"`
	Parse string `json:"parse"`
	IconEmoji string `json:"icon_emoji,omitempty"`
}

type Config struct {
	WebhookUrl string `json:"webhook_url"`
	Channel string `json:"channel"`
	Username string `json:"username"`
	IconEmoji string `json:"iconemoji"`
}

func ReadConfig() (*Config, error) {
	homeDir := ""
	usr, err := user.Current()
	if err == nil {
		homeDir = usr.HomeDir
	}
	log.Println("dir is :",homeDir,usr)

	for _, path := range []string{ homeDir + "/.slack-bot-golang.cfg", "./slack-bot-golang.cfg","/app/slack-bot-golang.cfg"} {
		log.Println("path:",path)
		file, err := os.Open(path)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			log.Println(err)
			return nil, err
		}
		log.Println("config file:",file)
		json.NewDecoder(file)
		conf := Config{}
		err = json.NewDecoder(file).Decode(&conf)
		if err != nil {
			log.Println("config json decode",&conf,err)
			return nil, err
		}
		log.Println("ReadConfig cfg:",conf)
		return &conf, nil
	}
	log.Println("don't read config file")

	return nil, errors.New("Config file not found")
}

func SlackPost(imegeUrl string ,cfg *Config) error{

	var data SlackMsg
	data.Channel = cfg.Channel
	data.Username = cfg.Username
	data.IconEmoji = cfg.IconEmoji
	data.Parse = "full"
	data.Text = imegeUrl
	jsonData,err := json.Marshal(data)
	WebhookUrl := cfg.WebhookUrl
	
	log.Println("jsonData:",string(jsonData),err)

	resp, err := http.PostForm(WebhookUrl, url.Values{"payload": {string(jsonData)}})
	if err != nil{
		log.Println("post Form Error:",resp,err)
	}
	log.Println(resp.Status,resp.Body," : ",err)
	return err
}

