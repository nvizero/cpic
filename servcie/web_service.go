package service

import (
	"context"
	bot "cpic/bot"
	db "cpic/db/sqlc"
	"cpic/model"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

var sbot bot.Crawler
var baseUrl string = "http://51sex.vip"

func Handle() []model.Sex51 {
	var collect []model.Sex51
	var bases = []string{
		"https://17sex.vip/Category/4584.html",
		"http://51sex.vip",
	}
	for _, url := range bases {
		col := FetchMain(url)
		collect = append(collect, col...)
	}
	return collect
}

func FetchMain(url string) []model.Sex51 {
	sbot.SetUrl(url)
	sbot.Fetch()
	sbot.JContent()
	if strings.Contains(url, "51sex") {
		return bot.Parse51SexIndex(sbot.Body, "a")
	} else {
		return bot.Parse17SexIndex(sbot.Body, "a")
	}
}

func FetchDoc(url string, style string) []string {
	sbot.SetUrl(url)
	sbot.Fetch()
	sbot.JContent()
	if strings.Contains(style, "51sex") {
		sbot.Type = "http://51sex.vip"
	} else {
		sbot.Type = "https://17sex.vip"
	}
	return bot.Get17SexContent(sbot.Body)
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
				Dt:        sql.NullString{String: row.Date, Valid: true},
				State:     sql.NullBool{Bool: false, Valid: true},
				Img:       row.Img,
				CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
			}
			_, err := store.CreatePost(context.Background(), arg)
			if err != nil {
				fmt.Println("err -- ", err)
			}

		}
	}
}
