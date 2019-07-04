package contollerTest

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

type User struct {
	Id         uint       `gorm:"primary_key"`
	Username   string
	Password   string
	State      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
	UserDetail UserDetail
	Posts      []Posts
}

type UserDetail struct {
	Id        uint       `gorm:"primary_key"`
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
	Id        uint       `gorm:"primary_key"`
	UserId    uint
	Title     string
	Body      string
	Link      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func Migrate(c *gin.Context) {

	db, err := gorm.Open("mysql", "homestead:123123@tcp(192.168.10.10:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	//db.AutoMigrate(&Hello{})

	if !db.HasTable("user") {
		db.Table("user").CreateTable(&User{})

	}
	if !db.HasTable("user_detail") {
		db.Table("user_detail").CreateTable(&UserDetail{})

	}
	if !db.HasTable("post") {
		db.Set("gorm:table_options", "ENGINE=InnoDB").Table("post").CreateTable(&Posts{})

	}

	//roll back
	_ = func(){
		db.DropTableIfExists("user")
	}
	_ = func() {
		db.Model(&User{}).ModifyColumn("state", "status")
		db.Model(&User{}).DropColumn("state")
		db.Model(&User{}).AddIndex("index_for_state","state")
		db.Model(&User{}).AddIndex("index_user_for_multi","state","username")
		db.Model(&User{}).AddUniqueIndex("index_unique_state","state")
		db.Model(&User{}).RemoveIndex("index_unique_state")

		//原生sql 执行
		sql := "select * from user where 1"
		db.Exec(sql)
	}

	c.JSON(200, gin.H{
		"data": true,
	})
}

func Create(c *gin.Context) {
	username := c.DefaultQuery("key", "migrate username")
	password := c.DefaultQuery("secret", "migrate password")

	nickname := c.DefaultQuery("nickname", "joen")
	sex, _ := strconv.Atoi(c.DefaultQuery("sex", "1"))
	avatar := c.DefaultQuery("avatar", "http://image.com")
	profile := c.DefaultQuery("profile", "http://homepage.com")

	db, err := gorm.Open("mysql", "homestead:123123@tcp(192.168.10.10:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.SingularTable(true)

	test := User{Username: username, Password: password}

	db.Create(&test)

	userDetail := UserDetail{NickName: nickname, Sex: sex, Avatar: avatar, Profile: profile, UserId: test.Id}

	db.Create(&userDetail)

	c.JSON(200, gin.H{
		"data": test,
	})
}

func Search(c *gin.Context) {

	page, limit, offset, keyword := listInit(c)

	db, err := gorm.Open("mysql", "homestead:123123@tcp(192.168.10.10:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.SingularTable(true)

	var Data []User
	var Count int

	db.Debug().Where("username like ?", "%"+keyword+"%").Limit(limit).Offset(offset).Order("id desc").Find(&Data) // 查询code为l1212的Test
	db.Debug().Model(&Data).Where("username like ?", "%"+keyword+"%").Count(&Count)                               // 查询code为l1212的Test

	c.JSON(200, gin.H{
		"data":  Data,
		"count": Count,
		"page":  page,
		"limit": limit,
	})

}

func First(c *gin.Context) {
	id := c.DefaultQuery("id", "1")
	db, err := gorm.Open("mysql", "homestead:123123@tcp(192.168.10.10:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var Data User

	db.Debug().Where("id = ? ", id).First(&Data)

	c.JSON(200, gin.H{
		"data": Data,
	})

}

func Update(c *gin.Context) {
	id := c.Query("id")
	db, err := gorm.Open("mysql", "homestead:123123@tcp(192.168.10.10:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	//var Hello Hello

	var first User

	db.Debug().Where("id = ? ", id).First(&first)

	db.Debug().Model(&first).Update(map[string]interface{}{"Key": "testKey", "Secret": "testSecret"})

	c.JSON(200, gin.H{
		"data": first,
	})

}

func Delete(c *gin.Context) {
	id := c.Query("id")
	db, err := gorm.Open("mysql", "homestead:123123@tcp(192.168.10.10:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	//var Hello Hello

	var first User

	db.Debug().Where("id = ? ", id).First(&first)

	db.Debug().Delete(&first)

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
