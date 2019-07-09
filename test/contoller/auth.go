package contoller

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Login(c *gin.Context) {

	session := sessions.Default(c)

	user_id := session.Get("id")

	fmt.Println(user_id)

	if user_id != nil{
		c.Redirect(302,"/index")
	}

	if strings.ToLower(c.Request.Method) == "get" {

		c.Abort()
		c.HTML(http.StatusOK, "login.html", gin.H{
			"Title": "作品欣赏",
		})

		return
	}

	username := c.PostForm("username")
	password := c.PostForm("password")

	var data User

	Orm.Debug().Where("username = ? ", username).Where("password = ? ", password).First(&data)


	if data.Id > 0 {
		session.Set("id", data.Id)
		_ = session.Save()

		c.Redirect(302, "/index")
	} else {
		c.Redirect(302, "/login")
	}
}

func Logout(c *gin.Context) {

	session := sessions.Default(c)

	session.Clear()
	_ = session.Save()

	c.Redirect(http.StatusFound,"/login")

}

func Register(c *gin.Context) {

	username := c.DefaultQuery("username", "admin")
	password := c.DefaultQuery("password", "admin")

	session := sessions.Default(c)

	user := User{Username: username, Password: password}

	Orm.Create(&user)

	session.Set("id", user.Id)
	_ = session.Save()

	c.JSON(200, gin.H{
		"data": user,
	})

}
