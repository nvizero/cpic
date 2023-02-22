package bot

import (
	"archive/zip"
	"bytes"
	"cpic/model"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
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

var imgBase string = "https://image.xjqh1007.com/tw0625"

// 取內容
func Get17SexContent(body []byte, id string) []string {
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
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		src, _ := s.Attr("src")
		imgName := src[strings.LastIndex(src, "/")+1:]
		if len(strings.Split(src, "/")) > 4 {
			if strings.Contains(src, "http") {
				downloadImg(src, imgName, id)
			} else {
				downloadImg("http:"+src, imgName, id)
			}
		}
		s.SetAttr("src", imgBase+"/"+id+"/"+imgName)
	})
	zipDir(id)
	sexCont.Content = doc.Find(dom).Text()
	html, _ := doc.Find(dom).Html()
	s := []string{title, time, html}
	//return doc.Find(dom).Html()
	return s
}
func Get51SexContent(body []byte, id string) []string {
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
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		src, _ := s.Attr("src")
		imgName := src[strings.LastIndex(src, "/")+1:]
		if len(strings.Split(src, "/")) > 4 {
			if strings.Contains(src, "http") {
				downloadImg(src, imgName, id)
			} else {
				downloadImg("http:"+src, imgName, id)
			}
		}
		s.SetAttr("src", imgBase+"/"+id+"/"+imgName)
	})
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

func downloadImg(imageUrl, imgName, dirName string) {
	response, err := http.Get(imageUrl)
	if err != nil {
		fmt.Println("Error while downloading image:", err)
		return
	}
	defer response.Body.Close()
	dirPath := "imgs"
	// 組合資料夾路徑
	dirFullPath := filepath.Join(dirPath, dirName)

	// 檢查資料夾是否存在
	if _, err := os.Stat(dirFullPath); os.IsNotExist(err) {
		// 資料夾不存在，建立資料夾
		if err := os.MkdirAll(dirFullPath, 0755); err != nil {
			log.Fatal(err)
		}
	} else if err != nil {
		// 其他錯誤
		log.Fatal(err)
	}
	file, err := os.Create(dirPath + "/" + dirName + "/" + imgName)
	if err != nil {
		fmt.Println("Error while creating file:", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Println("Error while saving image:", err)
		return
	}
}

func zipDir(zipName string) {
	dirp, _ := os.Getwd()
	dir := dirp + "/imgs/" + zipName
	zipFile, err := os.Create(zipName + ".zip") // 建立壓縮檔案
	if err != nil {
		panic(err)
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	// 走訪目錄下的檔案和子目錄
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 創建檔案或目錄的標頭
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = path

		// 如果是目錄，只創建標頭即可
		if info.IsDir() {
			header.Name += "/"
			header.Method = zip.Store // 目錄不需要壓縮，設定為存儲模式
			_, err = archive.CreateHeader(header)
			return err
		}

		// 建立檔案的標頭並寫入內容
		header.Method = zip.Deflate // 壓縮模式
		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(writer, file)
		return err
	})
	if err != nil {
		panic(err)
	}
}
