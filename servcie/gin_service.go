package service

import (
	db "cpic/db/sqlc"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.LoadHTMLGlob("view/*.tmpl")
	router.Static("/assetPath", "./asset")
	router.Use(cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/", func(c *gin.Context) {
		sex51links := WebSeseav()
		//for _, row := range sex51links {
		//	arg := db.CreatePostParams{
		//		Title: sql.NullString{String: row.Title, Valid: true},
		//		Link:  sql.NullString{String: row.Link, Valid: true},
		//		Img:   row.Img,
		//	}
		//	server.store.CreatePost(context.Background(), arg)
		//}
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"datas": sex51links,
			"baseu": "http://51sex.vip/",
		})
	})
	server.router = router
	return server
}

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
		//for _, row := range sex51links {
		//	arg := db.CreatePostParams{
		//		Title: sql.NullString{String: row.Title, Valid: true},
		//		Link:  sql.NullString{String: row.Link, Valid: true},
		//		Img:   row.Img,
		//	}
		//	server.store.CreatePost(context.Background(), arg)
		//}
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

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
