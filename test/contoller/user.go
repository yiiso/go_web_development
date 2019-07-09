package contoller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	_ "test/models"
	mysql "test/models"
	"time"
)

type User struct {
	Id         uint `gorm:"primary_key"`
	Username   string
	Password   string
	State      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
	UserDetail UserDetail `gorm:"ForeignKey:UserId"`
	Posts      []Posts
}

type UserDetail struct {
	Id        uint `gorm:"primary_key"`
	UserId    uint
	NickName  string
	Sex       int
	Avatar    string
	Profile   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type Posts struct {
	Id        uint `gorm:"primary_key"`
	UserId    uint
	Title     string
	Body      string
	Link      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

var Orm = mysql.GetGorm()


func Migrate(c *gin.Context) {

	//Db.AutoMigrate(&Hello{})

	if !Orm.HasTable("user") {
		Orm.Table("user").CreateTable(&User{})

	}
	if !Orm.HasTable("user_detail") {
		Orm.Table("user_detail").CreateTable(&UserDetail{})

	}
	if !Orm.HasTable("post") {
		Orm.Set("gorm:table_options", "ENGINE=InnoDB").Table("post").CreateTable(&Posts{})

	}

	//roll back
	_ = func() {
		Orm.DropTableIfExists("user")
	}
	_ = func() {
		Orm.Model(&User{}).ModifyColumn("state", "status")
		Orm.Model(&User{}).DropColumn("state")
		Orm.Model(&User{}).AddIndex("index_for_state", "state")
		Orm.Model(&User{}).AddIndex("index_user_for_multi", "state", "username")
		Orm.Model(&User{}).AddUniqueIndex("index_unique_state", "state")
		Orm.Model(&User{}).RemoveIndex("index_unique_state")

		//原生sql 执行
		sql := "select * from user where 1"
		Orm.Exec(sql)
	}

	c.JSON(200, gin.H{
		"data": true,
	})
}


func Index(c *gin.Context){
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title": "作品欣赏",
	})
}


func Create(c *gin.Context) {
	username := c.DefaultQuery("key", "migrate username")
	password := c.DefaultQuery("secret", "migrate password")

	nickname := c.DefaultQuery("nickname", "joen")
	sex, _ := strconv.Atoi(c.DefaultQuery("sex", "1"))
	avatar := c.DefaultQuery("avatar", "http://image.com")
	profile := c.DefaultQuery("profile", "http://homepage.com")

	test := User{Username: username, Password: password}

	Orm.Create(&test)

	userDetail := UserDetail{NickName: nickname, Sex: sex, Avatar: avatar, Profile: profile, UserId: test.Id}

	Orm.Create(&userDetail)

	c.JSON(200, gin.H{
		"data": test,
	})
}

func Search(c *gin.Context) {

	page, limit, offset, keyword := listInit(c)

	var Data []User
	var Count int

	Orm.Debug().Where("username like ?", "%"+keyword+"%").Limit(limit).Offset(offset).Order("id desc").Find(&Data) // 查询code为l1212的Test
	Orm.Debug().Model(&Data).Where("username like ?", "%"+keyword+"%").Count(&Count)                               // 查询code为l1212的Test

	c.JSON(200, gin.H{
		"data":  Data,
		"count": Count,
		"page":  page,
		"limit": limit,
	})

}

func First(c *gin.Context) {
	id := c.DefaultQuery("id", "1")

	var Data User

	Orm.Debug().Where("id = ? ", id).Preload("UserDetail","UserDetail").First(&Data)

	c.JSON(200, gin.H{
		"data": Data,
	})

}

func Update(c *gin.Context) {
	id := c.Query("id")

	//var Hello Hello

	var first User

	Orm.Debug().Where("id = ? ", id).First(&first)

	Orm.Debug().Model(&first).Update(map[string]interface{}{"Key": "testKey", "Secret": "testSecret"})

	c.JSON(200, gin.H{
		"data": first,
	})

}

func Delete(c *gin.Context) {
	id := c.Query("id")
	//var Hello Hello

	var first User

	Orm.Debug().Where("id = ? ", id).First(&first)

	Orm.Debug().Delete(&first)

	c.JSON(200, gin.H{
		"data": true,
	})

}

func listInit(c *gin.Context) (int, int, int, string) {

	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))

	keyword := c.DefaultQuery("keyword", "")

	if (err != nil) {
		panic("数据转换出错")
	}

	offset := page * limit

	return page, limit, offset, keyword

}
