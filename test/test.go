package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"test/contollerTest"
)

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/migrate", contollerTest.Migrate)
	r.GET("/ping", contollerTest.Search)
	r.GET("/create", contollerTest.Create)
	r.GET("/first", contollerTest.First)
	r.GET("/update", contollerTest.Update)
	r.GET("/list", contollerTest.Search)
	r.GET("/delete", contollerTest.Delete)

	r.Run(":8989") // listen and serve on 0.0.0.0:8080
}
