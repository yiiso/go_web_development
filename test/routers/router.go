package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	"os"
	"test/MiddleWare"
	user "test/contoller"
	"time"
)

func Register() *gin.Engine {

	/*	gin.DisableConsoleColor()//禁用请求日志控制台字体颜色

		r := gin.New()
		r.Use(gin.Logger(),gin.Recovery())*/

	gin.DisableConsoleColor() //保存到文件不需要颜色
	file, _ := os.Create("logs/" + time.Now().Format("2006-01-02") + ".log")
	gin.DefaultWriter = file
	//gin.DefaultWriter = io.MultiWriter(file) 效果是一样的
	r := gin.Default()

	var cfg *ini.File
	var iniE error
	cfg, iniE = ini.Load("conf/database.ini", "conf/app.ini")
	if iniE != nil {
		panic("config don't reader import!")
	}

	mode := cfg.Section("").Key("app_mode").String()

	host := cfg.Section(mode).Key("redis.host").String()
	// 端口
	port := cfg.Section(mode).Key("redis.port").String()

	store, err := redis.NewStore(10, "tcp", host+":"+port, "", []byte("secret"))

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
		v1.GET("migrate", user.Migrate)
		v1.GET("wx", user.WxApi)

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
