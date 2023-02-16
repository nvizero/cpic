package service

import (
	"html/template"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routes() {

	r := gin.Default()
	r.LoadHTMLGlob("view/*.tmpl")
	r.Static("/assetPath", "./asset")

	r.Use(cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/", func(c *gin.Context) {
		sex51links := WebSeseav()
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"datas": sex51links,
			"baseu": "http://51sex.vip/",
		})
	})

	r.GET("/doc", func(c *gin.Context) {
		id := c.Query("id")
		arys := FetchDoc(id)
		c.HTML(http.StatusOK, "doc.tmpl", gin.H{
			"content": template.HTML(arys[2]),
			"title":   arys[0],
			"time":    arys[1],
		})
	})

	r.Run(":8333") // listen and serve on 0.0.0.0:8080
}
