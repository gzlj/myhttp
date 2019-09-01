package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gzlj/myhttp/pkg/db"
	"strconv"
	//"github.com/swaggo/gin-swagger"
)

func init() {

}

func main() {
	server := &APIServer{
		engine: gin.Default(),
	}
	server.registryApi()
	server.engine.Run(":8080")
}

type APIServer struct {
	engine *gin.Engine
}

func (s *APIServer) registryApi() {
	registryHello(s.engine)
	registryUser(s.engine)
	registryLike(s.engine)
}

func (s *APIServer) Start() {

}

func registryHello(engine *gin.Engine) {
	engine.GET("/hello", func(ctx *gin.Context) {
		message := ctx.Query("message")
		fmt.Print("wo kao", message)
		ctx.JSON(200, gin.H{
			"message": message,
		})
	})
}

func registryUser(r *gin.Engine) {
	r.GET("/user/:name", func(ctx *gin.Context) {
		userName := ctx.Param("name")
		ctx.JSON(200, gin.H{
			"user": userName,
		})
	})

	r.GET("/user", func(context *gin.Context) {
		var user User
		if err := context.ShouldBind(&user); err != nil {
			context.JSON(500, gin.H{
				"error": err,
			})
			return
		}
		context.JSON(200, user)
	})
}

type User struct {
	Name string `form:"name"`
	Age string `form:"age"`
}

func registryLike(r *gin.Engine) {
	r.GET("/like/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		var likeId int
		var err error
		if likeId, err = strconv.Atoi(id); err != nil {
			ctx.JSON(400, "input id is not correct.")
			return
		}
		fmt.Print("likeId: ", likeId)
		var like *db.Like
		if like = db.QueryById(likeId); like == nil {
			ctx.JSON(404, "Not found")
			return
		}
		ctx.JSON(200, like.ToDto())
	})

	r.POST("/like", func(c *gin.Context) {
		var dto db.LikeDto
		var err error
		if err = c.ShouldBindJSON(&dto); err != nil {
			c.JSON(400, "requet body is not correct.")
			return
		}
		like := dto.ToLike()
		if err = db.ADDLike(&like); err != nil {
			c.JSON(500, "cannot create like in db.")
		}
		c.JSON(200, "OK")
	})
}