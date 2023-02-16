package service

import (
	bot "cpic/bot"
	"cpic/model"
	"fmt"
)

var sbot bot.Crawler
var baseUrl string = "http://51sex.vip"

func FetchDoc(doc string) []string {
	url := baseUrl + doc
	sbot.SetUrl(url)
	sbot.Fetch()
	sbot.JContent()
	htmlcc := bot.GetContent(sbot.Body)
	return htmlcc
}

func WebSeseav() []model.Sex51 {
	sbot.SetUrl(baseUrl)
	sbot.Fetch()
	sbot.JContent()
	alinks := bot.GetIndex(sbot.Body, "a")
	fmt.Println("--", alinks)
	return alinks
}
