package setting

import (
	"log"

	"github.com/go-ini/ini"
)

var (
	Cfg      *ini.File
	RunMode  string
	HttpPort int
	PageSize int
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, _ := Cfg.GetSection("server")
	HttpPort = sec.Key("HTTP_PORT").MustInt(2333)
}

func LoadApp() {
	sec, _ := Cfg.GetSection("app")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
