package service

import (
	"context"
	bot "cpic/bot"
	db "cpic/db/sqlc"
	"cpic/model"
	"database/sql"
	"fmt"
	"time"
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

func NewDBServer(store db.Store) []model.Sex51 {
	sbot.SetUrl(baseUrl)
	sbot.Fetch()
	sbot.JContent()
	return bot.GetIndex(sbot.Body, "a")
}

// 跟資料庫比較 爬取的小說
func CheckDataAndInsert(store db.Store, craDatas []model.Sex51) {
	posts, _ := store.GetPosts(context.Background())
	//迴圈比對兩個陣列的元素
	titleMap := make(map[string]bool)
	for _, row := range posts {
		titleMap[row.Title.String] = false
	}

	for _, row := range craDatas {
		// 若 title 在 map 中出現過，則找到相同的 title
		_, ok := titleMap[row.Title]
		if !ok {
			//b.EnableNovelLinks = append(b.EnableNovelLinks, link)
			arg := db.CreatePostParams{
				Title:     sql.NullString{String: row.Title, Valid: true},
				Link:      sql.NullString{String: row.Link, Valid: true},
				State:     sql.NullBool{Bool: false, Valid: true},
				Img:       row.Img,
				CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
			}
			_, err := store.CreatePost(context.Background(), arg)
			if err != nil {
				fmt.Println("err -??-- ", err)
			}

		}
	}
}
