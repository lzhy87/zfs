package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
	"zfs/utils"
)

var str string
var timeout int

func main() {

	flag.StringVar(&str, "m", "", "邮箱地址")
	flag.IntVar(&timeout, "t", 30, "监控健康间隔时间")
	flag.Parse()
	switch {
	case str != "" && timeout >= 30:
		monitor()
	default:
		fmt.Println("输入参数错误！")
	}

}
func monitor() {
	hostname, _ := os.Hostname()
	for {
		time.Sleep(time.Second * time.Duration(timeout))
		cmd := exec.Command("/bin/bash", "-c", `/sbin/zpool status | grep state:| awk '{print $2}'`)
		cmd.Stdin = strings.NewReader("some input")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(out.String())
		if out.String() == "DEGRADED\n" {
			utils.SendMail([]string{str}, "zfs sever was broken", hostname+" zfs服务器故障")
			fmt.Println("邮件已发送！")
			return
		}
	}
}
