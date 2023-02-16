package model

type Link struct {
	Href       string
	Title      string
	Txt        string
	Id         string
	NovelId    string
	NovelTitle string
}

type Novel struct {
	Url       string
	Title     string
	PartLinks []Link
	AllLinks  []Link
	Content   string
}

type Comp struct {
	Id    string
	Title string
}

type Sex51 struct {
	Title string
	Img   string
	Link  string
	Date  string
}

type Sex51Cont struct {
	Title   string
	Img     string
	Link    string
	Date    string
	Content string
}
