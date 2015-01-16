package main

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

	for _, path := range []string{ homeDir + "/.slack-bot-golang.cfg", "./slack-bot-golang.cfg"} {
		file, err := os.Open(path)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			log.Println(err)
			return nil, err
		}

		json.NewDecoder(file)
		conf := Config{}
		err = json.NewDecoder(file).Decode(&conf)
		if err != nil {
			log.Println("config json decode",&conf,err)
			return nil, err
		}
		return &conf, nil
	}

	return nil, errors.New("Config file not found")
}

func SlackPost(cfg *Config) error{

	var data SlackMsg
	data.Channel = cfg.Channel
	data.Username = cfg.Username
	data.IconEmoji = cfg.IconEmoji
	data.Parse = "full"
	data.Text = "ちょっとソースを書き換えたもので"
	jsonData,err := json.Marshal(data)
	log.Println(string(jsonData),err)

	WebhookUrl := cfg.WebhookUrl

	//log.Println(data,string(jsonData))

	resp, err := http.PostForm(WebhookUrl, url.Values{"payload": {string(jsonData)}})
log.Println(resp.Status,resp.Body," : ",err)
return err
}

