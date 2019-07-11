package models

import (
	"fmt"
	"github.com/go-ini/ini"
	"os"
)

var cfg * ini.File

func init() {
	var err error

	// load配置
	cfg, err = ini.Load("conf/database.ini", "conf/app.ini","conf/wechat.ini")
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
}

func GetCfg() * ini.File {
	return cfg
}