package sql

import (
	"cpic/model"
	"strconv"
)

var ExistNovels, ExistParts []model.Link

func GetAllParts() []model.Link {
	var id, title string
	sqltxt := "select id ,title from novel_parts"
	rows := Query(sqltxt)
	for rows.Next() {
		rows.Scan(&id, &title)
		ExistParts = append(ExistParts, model.Link{
			Id:    id,
			Title: title,
		})
	}
	return ExistParts
}

type Cont struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func GetNovel(id int) Cont {
	var content, title string
	var nvl = model.Novel{}
	sqltxt := "select content ,title from novel_part_contents where id = " + strconv.Itoa(id) + " order by id asc"
	rows := Query(sqltxt)
	for rows.Next() {
		rows.Scan(&content, &title)
	}
	nvl.Content = content
	nvl.Title = title
	var cont = Cont{
		Title:   title,
		Content: content,
	}
	return cont
}

// 每一個章節
func GetNovelParts(novel_id int) []model.Link {
	var id, title string
	var NovelParts []model.Link
	//sqltxt := "select id,title from novel_parts where novel_id = '" + strconv.Itoa(novel_id) + "' order by id desc"

	sqltxt := "select c.id as id, a.title as title from novel_parts a  left join    novel_part_contents c on a.novel_id = c.novel_part_id  where a.novel_id = '" + strconv.Itoa(novel_id) + "'  order by a.id desc ;"
	rws := Query(sqltxt)
	for rws.Next() {
		rws.Scan(&id, &title)
		NovelParts = append(NovelParts, model.Link{
			Id:    id,
			Title: title,
		})
	}
	return NovelParts
}

func GetAllNovels() []model.Link {
	var id, title string
	sqltxt := "select id ,title from novels"
	rows := Query(sqltxt)
	for rows.Next() {
		rows.Scan(&id, &title)
		ExistNovels = append(ExistNovels, model.Link{
			Id:    id,
			Title: title,
		})
	}
	return ExistNovels
}
