package service

import (
	"context"
	bot "cpic/bot"
	db "cpic/db/sqlc"
	"cpic/model"
	"database/sql"
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
	return alinks
}

func NewDBServer(store db.Store) {
	sbot.SetUrl(baseUrl)
	sbot.Fetch()
	sbot.JContent()
	alinks := bot.GetIndex(sbot.Body, "a")
	//store.CreatePost
	for _, row := range alinks {
		arg := db.CreatePostParams{
			Title: sql.NullString{String: row.Title, Valid: true},
			Link:  sql.NullString{String: row.Link, Valid: true},
			State: sql.NullBool{Bool: false, Valid: true},
			Img:   row.Img,
		}
		_, err := store.CreatePost(context.Background(), arg)
		if err != nil {
			fmt.Println("err --- ", err)
		}
	}
}
