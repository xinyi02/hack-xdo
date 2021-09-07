package models

import (
	"errors"
	"os"
	"os/exec"
	"regexp"
	//"strings"

	"github.com/beego/beego/v2/client/httplib"
	"github.com/beego/beego/v2/core/logs"
)

var version = "2021090710"
var describe = "日常更新"
var AppName = "xdd"
var pname = regexp.MustCompile(`/([^/\s]+)`).FindStringSubmatch(os.Args[0])[1]

func initVersion() {
	if Config.Version != "" {
		version = Config.Version
	}
	logs.Info("检查更新" + version)
	value, err := httplib.Get(GhProxy + "https://raw.githubusercontent.com/xnyi02/hack-xdo/main/models/version.go").String()
	if err != nil {
		logs.Info("更新版本的失败")
	} else {
		// name := AppName + "_" + runtime.GOOS + "_" + runtime.GOARCH
		if match := regexp.MustCompile(`var version = "(\d{10})"`).FindStringSubmatch(value); len(match) != 0 {
			des := regexp.MustCompile(`var describe = "([^"]+)"`).FindStringSubmatch(value)
			if len(des) != 0 {
				describe = des[1]
			}
			if match[1] > version {
				err := Update(&Sender{})
				if err != nil {
					logs.Warn("更新失败,", err)
					return
				}
				(&JdCookie{}).Push("苏青更新：" + describe)
				Daemon()
			}
		}
	}
}

func Update(sender *Sender) error {
	//sender.Reply("苏青开始拉取代码")
	//rtn, err := exec.Command("sh", "-c", "cd "+ExecPath+" && git stash && git pull").Output()
	//if err != nil {
	//	return errors.New("苏青拉取代失败：" + err.Error())
	//}
	//t := string(rtn)
	//if !strings.Contains(t, "changed") {
	//	if strings.Contains(t, "Already") || strings.Contains(t, "已经是最新") {
	//		return errors.New("苏青已是最新版啦")
	//	} else {
	//		return errors.New("苏青拉取代失败：" + t)
	//	}
	//} else {
	//	sender.Reply("苏青拉取代码成功")
	//}

	sender.Reply("苏青正在编译程序")
	_, err := exec.Command("sh", "-c", "cd "+ExecPath+" && go build -o "+pname).Output()
	if err != nil {
		return errors.New("苏青编译失败：" + err.Error())
	} else {
		sender.Reply("苏青编译成功")
	}
	return nil
}
