package MiddleWare

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context){

	//2.若session存在，则继续进行请求操作，并将session的有效时间重新设置一次；
	//3.若session不存在，则判断cookie是否存在?
	//4.若cookie存在，使用该cookie完成自动登录;

	session := sessions.Default(c)
	id := session.Get("id")

	fmt.Println(id)
	if id == nil {
		c.Abort()
		c.Redirect(302,"/login")
		return
	}
	session.Set("id", id)
	_ = session.Save()

/*	var User contoller.User

	contoller.Orm.Debug().Where("id = ? ", id).Preload("UserDetail","UserDetail").First(&User)
	c.JSON(200, gin.H{"data": User})*/

}