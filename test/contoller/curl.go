package contoller

//抓取页面

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mikemintang/go-curl"
	"net/http"
)

func CurlExe(url string, headers map[string]string, data map[string]interface{}) string {
	req := curl.NewRequest()
	resp, err := req.
		SetUrl(url).
		SetHeaders(headers).
		SetPostData(data).
		Post()

	if err != nil {
		fmt.Println(err)
		return "{}"
	} else {
		if resp.IsOk() {
			return (resp.Body)
		} else {
			return "{}"
		}

	}

}

func GetCountry(c *gin.Context) {

	url := "https://img2.keketour.com/country.json"
	headers := map[string]string{}
	data := map[string]interface{}{}
	res := CurlExe(url, headers, data)

	fmt.Println(res)

	var jsonData interface{}

	_ = json.Unmarshal([]byte(res), &jsonData)

	c.JSON(http.StatusOK, gin.H{
		"data": jsonData,
	})
}

func GetGuideCard(c *gin.Context)  {

	url := "http://jianguan.12301.cn/data/guide/verify"
	headers := map[string]string{
		"Accept":"application/json, text/javascript, */*; q=0.01",
		"Content-Type":"application/json",

	}
	data := map[string]interface{}{
		"type":3,
		"value":"JDQ5945O",
	}
	res := CurlExe(url, headers, data)

	var jsonData interface{}

	_ = json.Unmarshal([]byte(res), &jsonData)

	c.JSON(http.StatusOK, gin.H{
		"data": jsonData,
	})
}
