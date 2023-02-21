package bot

import (
	"bytes"
	"cpic/model"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func PureContent(body []byte) string {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	var content string
	content = doc.Text()
	return content
}

// 取內容
func Get17SexContent(body []byte) []string {
	var sexCont model.Sex51Cont
	dom := ".artcontent"
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	title := doc.Find("h1.cps-post-title.entry-title").Text()
	time := doc.Find(".entry-date").Text()
	doc.Find(".news-item").Remove()
	doc.Find("style").Remove()
	doc.Find("script").Remove()
	sexCont.Content = doc.Find(dom).Text()
	html, _ := doc.Find(dom).Html()
	s := []string{title, time, html}
	//return doc.Find(dom).Html()
	return s
}
func Get51SexContent(body []byte) []string {
	var sexCont model.Sex51Cont
	dom := ".headling_wrod_main_box_edit"
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	title := doc.Find(".headling_word_main_box_title").Text()
	time := doc.Find(".headling_word_main_box_time").Text()
	doc.Find(".news-item").Remove()
	doc.Find("style").Remove()
	doc.Find("script").Remove()
	doc.Find("#compass-fit-4302731").Remove()
	sexCont.Content = doc.Find(dom).Text()
	html, _ := doc.Find(dom).Html()
	s := []string{title, time, html}
	//return doc.Find(dom).Html()
	return s
}

// 取每一個dom
func Parse51SexIndex(body []byte, dom string) []model.Sex51 {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	var sex51links []model.Sex51
	var slink model.Sex51
	doc.Find(dom).Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		slink.Link = strings.TrimSpace(link)
		s.Find(".headling_main_box_time").Each(func(i1 int, st *goquery.Selection) {
			slink.Date = strings.TrimSpace(st.Text())
		})
		s.Find(".headling_main_box_title").Each(func(i1 int, s2 *goquery.Selection) {
			slink.Title = strings.TrimSpace(s2.Text())
		})
		s.Find("img").Each(func(i2 int, s3 *goquery.Selection) {
			link, _ := s3.Attr("src")
			slink.Img = strings.TrimSpace(link)
			//fmt.Println("img  : ", link)
		})
		if len(slink.Title) > 1 && rmLink(link) {
			sex51links = append(sex51links, slink)
		}
	})
	return sex51links
}

func Parse17SexIndex(body []byte, dom string) []model.Sex51 {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	var sex51links []model.Sex51
	var slink model.Sex51
	doc.Find(dom).Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		slink.Link = strings.TrimSpace(link)
		s.Find(".post-list-date").Each(func(i1 int, st *goquery.Selection) {
			slink.Date = strings.TrimSpace(st.Text())
		})
		s.Find(".entry-title").Each(func(i1 int, s2 *goquery.Selection) {
			slink.Title = strings.TrimSpace(s2.Text())
		})
		s.Find("img").Each(func(i2 int, s3 *goquery.Selection) {
			link, _ := s3.Attr("src")
			slink.Img = strings.TrimSpace(link)
			//fmt.Println("img  : ", link)
		})
		if len(slink.Title) > 1 && rmLink(link) {
			sex51links = append(sex51links, slink)
		}
	})
	return sex51links
}

func rmLink(link string) bool {
	var outStrs = []string{"inks", "about", "email", "javascript"}
	var fj bool = true
	for _, str := range outStrs {
		if strings.Contains(link, str) {
			fj = false
		}
	}
	return fj

}
