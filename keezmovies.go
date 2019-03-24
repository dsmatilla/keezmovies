package keezmovies

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const apiURL = "http://www.keezmovies.com/wapi/"
const APITimeout = 5

func GetVideoByID(ID string) KeezmoviesSingleVideo {
	timeout := time.Duration(APITimeout * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, _ := client.Get(fmt.Sprintf(apiURL+"getVideoById?output=json&video_id=%s", ID))
	b, _ := ioutil.ReadAll(resp.Body)
	var result KeezmoviesSingleVideo
	err := json.Unmarshal(b, &result)
	if err != nil {
		log.Println(err)
	}
	return result

}

func GetVideoEmbedCode(ID string) KeezmoviesEmbedCode {
	timeout := time.Duration(APITimeout * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, _ := client.Get(fmt.Sprintf(apiURL+"getVideoEmbedCode?output=json&video_id=%s", ID))
	b, _ := ioutil.ReadAll(resp.Body)
	var result KeezmoviesEmbedCode
	err := json.Unmarshal(b, &result)
	if err != nil {
		log.Println(err)
	}
	return result
}
