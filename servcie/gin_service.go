package service

import (
	"context"
	db "cpic/db/sqlc"
	"database/sql"
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func Routes(store db.Store) {
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
		sex51links, _ := store.GetPosts(context.Background())
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"datas": sex51links,
			"baseu": "http://51sex.vip/",
		})
	})

	r.GET("/doc", func(c *gin.Context) {
		id := c.Query("id")
		doc := c.Query("doc")
		var arys []string
		if strings.Contains(doc, "17sex") {
			arys = FetchDoc(doc, "17sex", id)
		} else {
			arys = FetchDoc(doc, "51sex", id)
		}
		arg := db.UpdatePostParams{
			Link:    sql.NullString{String: id, Valid: true},
			State:   sql.NullBool{Bool: true, Valid: true},
			Content: sql.NullString{String: strings.TrimSpace(arys[2]), Valid: true},
		}
		store.UpdatePost(context.Background(), arg)
		c.HTML(http.StatusOK, "doc.tmpl", gin.H{
			"content": template.HTML(arys[2]),
			"title":   arys[0],
			"time":    arys[1],
		})
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
