package utils

import (
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

var (
	// mailTo []string
	// // //邮件主题
	// // subject = "zfs存储服务出现故障"
	// // // 邮件正文
	// // body string
	// // err  error

	// //配置文件内容
	hostname string

	// user   string
	// pass   string
	// host   string
	// port   string
	// mailto string
)

func init() {

	hostname, _ = os.Hostname()
	// conf, err := config.NewConfig("ini", "app.conf")
	// if err != nil {
	// 	log.Fatalf("config配置件读取失败%v", err)
	// }

	// user = conf.String("mail::user")
	// pass = conf.String("mail::pass")
	// host = conf.String("mail::host")
	// port = conf.String("mail::port")
	// mailto = conf.String("mail::mailto")
	// mailTo = strings.Fields(mailto)

	// if err != nil {
	// 	Logs.Error("获取hostname失败：", err)

	// }
}

func SendMail(mailTo []string, subject string, body string) error {
	//定义邮箱服务器连接信息，如果是阿里邮箱 pass填密码，qq邮箱填授权码
	mailConn := map[string]string{
		"user": "55900695@qq.com",
		"pass": "valhpmceiumbbibf",
		"host": "smtp.qq.com",
		"port": "465",
		// "user": user,
		// "pass": pass,
		// "host": host,
		// "port": port,
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()
	m.SetHeader("From", "zfs server "+hostname+" <"+mailConn["user"]+">") //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("To", mailTo...)                                          //发送给多个用户
	m.SetHeader("Subject", subject)                                       //设置邮件主题
	m.SetBody("text/html", body)                                          //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	return err

}
