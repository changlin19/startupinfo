package startupinfo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

var info = `
********************************************************
*	启动时间: %s
*	进程ID: %d
*	配置文件地址: %s
*	配置文件的信息: 
*    %s
*	最后一次修改时间: %s
*	最后一次修改内容: %s
********************************************************
`

func Print(configFilePath string, configInfo map[string]interface{}, lastTime, lastInfo string) {
	l, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(fmt.Sprintf("load Asia/Shanghai location error:%s", err.Error()))
	}
	cMsg, err := json.Marshal(configInfo)
	if err != nil {
		panic(fmt.Sprintf("marshal config info error:%s,data:%v", err.Error(), configInfo))
	}
	var out bytes.Buffer
	err = json.Indent(&out, cMsg, "*    ", "\t")
	if err != nil {
		panic(fmt.Sprintf("json indent error:%s", err.Error()))
	}
	fmt.Printf(info, time.Now().In(l).String(), os.Getpid(), configFilePath, out.String(), lastTime, lastInfo)
}
