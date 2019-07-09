package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"test/MiddleWare"
	user "test/contoller"
)

func Register() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	store, err := redis.NewStore(10, "tcp", "192.168.10.10:6379", "", []byte("secret"))

	if err != nil {
		panic("redis don't connection!")
	}

	r.Use(sessions.Sessions("mysession", store)) //使用中间使
	//r.Use(MiddleWare.Auth)

	//articles := new(controllers.Search)

	v1 := r.Group("/")
	{

		v1.Any("login", user.Login)
		v1.GET("logout", user.Logout)
		v1.GET("register", user.Register)
	}


	v2 := r.Group("/").Use(MiddleWare.Auth)
	{

		v2.GET("index", user.Index)
		v2.Any("file", user.File)
		v2.GET("user/search", user.Search)
		v2.GET("user/first", user.First)
		v2.GET("user/update", user.Update)
		v2.GET("user/delete", user.Delete)
		v2.GET("country", user.GetCountry)
		v2.GET("guide", user.GetGuideCard)
	}

	//defer articles.Db.Close()
	return r
}
