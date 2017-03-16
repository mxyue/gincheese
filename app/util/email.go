package util

import (
	"fmt"
	"gincheese/app/config"
	"math/rand"
	"net/smtp"
	"regexp"
	"time"
)

const (
	email_regular = "^[a-z0-9]+([._\\-]*[a-z0-9])*@([a-z0-9]+[-a-z0-9]*[a-z0-9]+.){1,63}[a-z0-9]+$"
)

func SendCode(email, site string) error {
	rand.Seed(int64(time.Now().Nanosecond()))
	code := fmt.Sprintf("%d", rand.Intn(9000)+1000)
	to := []string{email}
	msg := fmt.Sprintf("To: %s\r\n", email) +
		"Subject: 邮箱验证\r\n" +
		"\r\n" +
		fmt.Sprintf("请好好保管，验证码是：%s \r\n", code)
	CacheSet(site, email, code)
	err := Send(to, []byte(msg))
	return err
}

func Send(emails []string, content []byte) error {
	auth := smtp.PlainAuth("", config.Settings["email_add"], config.Settings["email_pw"], config.Settings["email_host"])
	err := smtp.SendMail(config.Settings["email_host"]+":25", auth, config.Settings["email_add"], emails, content)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func EmailRegex(email string) bool {
	reg := regexp.MustCompile(email_regular)
	return reg.MatchString(email)
}
