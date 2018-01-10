package main

import (
	"github.com/hpcloud/tail"
	"fmt"
	//"strings"
	//"reflect"
	"strings"
	"encoding/json"
	"net/http"
	"io"
	"os"
	"flag"
)

var logPath , downLoadDirPath string
func main() {
	flag.StringVar(&logPath ,"p" , "" ,"输出网易音乐的日志路径...")
	flag.StringVar(&downLoadDirPath ,"d" , "" ,"输入文件将要下载到哪..")
	flag.Parse()
	if logPath == "" {
		fmt.Println("请输入网易音乐的日志路径，参数为 '-p'\r")
		return
	}
	if downLoadDirPath == "" {
		fmt.Println("输入文件将要下载到的目录，参数为 '-d'\r")
		return
	}
	logPath = logPath + "/music.163.log"
	downLoadDirPath = downLoadDirPath + "/"

	file , err := tail.TailFile(logPath , tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	//读取日志文件
	for line := range file.Lines  {
		if strings.Contains(line.Text , "player._$load") {
			//查找json
			start := strings.LastIndex(line.Text , "{")
			songJsonStr := string(line.Text[start:len(line.Text)])
			//解析json
			var s  map[string]interface{}
			err := json.Unmarshal([]byte(songJsonStr) , &s)
			if err != nil {
				continue
			}

			var songName , musicUrl string
			songName = s["songName"].(string)
			musicUrl = s["musicurl"].(string)
			downLoad(songName , musicUrl)
			fmt.Println(songName , musicUrl)
		}

	}

}

/**
   下载文件
 */
func downLoad (fileName , fileUrl string) {
	fileName = downLoadDirPath + fileName

	fileRes , err := http.Get(fileUrl)

	if err != nil {
		fmt.Println(err)
		return;
	}

	file , err := os.Create(fileName + ".mp3")

	if err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(file , fileRes.Body)

}
