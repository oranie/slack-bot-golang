package lib

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

type ResponseData struct {
	ResponseData Top `json:responseData`
}

type Top struct {
	Cursor          interface{} `json:cursor`
	Results         []Result    `json:results`
	ResponseDetails interface{} `json:responseDetails`
	ResponseStatus  interface{} `json:responseStatus`
}

type Result struct {
	GsearchResultClass  string `json:gsearchResultClass`
	Content             string `json:content`
	ContentNoFormatting string `json:contentNoFormatting`
	Height              string `json:height`
	Width               string `json:width`
	ImageId             string `json:imageId`
	OriginalContextUrl  string `json:originalContextUrl`
	TbHeight            string `json:tbHeight`
	TbWidth             string `json:tbWidth`
	TbUrl               string `json:tbUrl`
	Title               string `json:title`
	TitleNoFormatting   string `json:titleNoFormatting`
	UnescapedUrl        string `json:unescapedUrl`
	Url                 string `json:url`
	VisibleUrl          string `json:visibleUrl`
}

func FetchImageUrl(s string) (string, error) {
	client := &http.Client{Timeout: time.Duration(10) * time.Second}
	req, err := http.NewRequest("GET", "http://ajax.googleapis.com/ajax/services/search/images?", nil)
	if err != nil {
		log.Println(err)
		return "", err
	}

	values := url.Values{}
	values.Add("q", s)
	values.Add("v", "1.0")
	values.Add("rsz", "8")
	values.Add("start", "1")
	values.Add("filter", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1500.52 Safari/537.36")
	req.URL.RawQuery = values.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	var responseData ResponseData
	b, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(b, &responseData)

	randUrl := responseData.ResponseData.Results[(rand.Intn(8))].UnescapedUrl
	log.Println(randUrl)

	return randUrl, err
}
