package contoller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func File(c *gin.Context)  {
	if strings.ToLower(c.Request.Method) == "get"{

		c.HTML(http.StatusOK,"file.html",gin.H{
			"title":"文件上传",
		})
	}

	file, err := c.FormFile("file")

	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	filename := file.Filename
	//目标文件

	dst := "static/upload/" + filename

	err = c.SaveUploadedFile(file, dst)
	if err != nil{

		c.String(http.StatusBadRequest, "Bad request")
		return

	}

	c.JSON(http.StatusOK,gin.H{
		"msg":"文件上传成功",
		"path" : dst,
	})

}
