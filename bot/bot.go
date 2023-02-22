package bot

import (
	"cpic/model"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Crawler struct {
	FetchUrl         string
	BaseUrl          string
	Name             string
	Content          string
	Type             string
	Body             []byte
	NovelLinks       []model.Link
	EnableNovelLinks []model.Link
	EnablePartLinks  []model.Link
	PartLinks        []model.Link
	AllLinks         []model.Link
	Queue            []model.Link
	Sync             sync.Mutex
}

// 爬網頁
func (c *Crawler) Fetch() *Crawler {
	fmt.Println("sleep 3 secs and start crawler ..." + c.FetchUrl)
	time.Sleep(time.Second * 2)
	resp, err := http.Get(c.FetchUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	c.Body = body
	return c
}

// 取內容
func (c *Crawler) JContent() *Crawler {
	content := PureContent(c.Body)
	c.Content = strings.Trim(content, "")
	return c
}
func (c *Crawler) Dequeue() model.Link {
	temp := c.Queue[0]
	c.Queue = c.Queue[1:]
	return temp
}

func (c *Crawler) Enqueue(s model.Link) {
	c.Queue = append(c.Queue, s)
}

func (c *Crawler) Top() model.Link {
	return c.Queue[0]
}

func (c *Crawler) Empty() bool {
	return len(c.Queue) == 0
}

// 設定URL
func (c *Crawler) SetUrl(url string) *Crawler {
	c.FetchUrl = url
	return c
}
