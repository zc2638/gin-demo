package mails

import (
	"github.com/go-gomail/gomail"
	"io"
	"mime/multipart"
)

const (
	host = "smtp.partner.outlook.cn"
	port = 587
)

type MailServer struct {
	From    string //发送邮箱
	Pwd     string //发送邮箱密码
	Host    string //邮箱smtp host
	Port    int    //邮箱smtp port
	To      string //接收邮箱
	Subject string //标题
	Body    string //内容
}

func (m *MailServer) Send(file *multipart.FileHeader) error {

	if m.Host == "" {
		m.Host = host
	}
	if m.Port == 0 {
		m.Port = port
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", m.From)
	msg.SetHeader("To", m.To)
	// 抄送
	//  msg.SetAddressHeader("Cc", "dan@example.com", "Dan")
	msg.SetHeader("Subject", m.Subject)
	msg.SetBody("text/html", m.Body)

	if file != nil {
		msg.Attach(file.Filename, gomail.SetCopyFunc(func(w io.Writer) error {
			f, _ := file.Open()
			if _, err := io.Copy(w, f); err != nil {
				return err
			}
			return nil
		}))
	}

	mailer := gomail.NewDialer(m.Host, m.Port, m.From, m.Pwd)
	return mailer.DialAndSend(msg)
}