package controllers

import (
	"fmt"
	"net/smtp"

	"github.com/jasonlvhit/gocron"
)

func (sc *ServerController) DailyReport() {
	s := gocron.NewScheduler()
	s.Every(1).Day().At("14:14:00").Do(sc.SendEmail)
	s.Start()
}

func (sc *ServerController) SendEmail() {
	from := "pdnsm080701@gmail.com"
	password := "yzvaovycjwarpfet"

	toFirstEmailAddress := "pdnsm080701@gmail.com"
	toSecondEmailAddress := "anhntpvcs@gmail.com"
	to := []string{toFirstEmailAddress, toSecondEmailAddress}

	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	countServer, countServerOn, countServerOff, upTimeAvg := sc.CheckStatus()

	subject := "Subject: Server Daily Report\n"
	body := fmt.Sprintf("Total Server: %d\nTotal Server On: %d\nTotal Server Off: %d\nUptime Average: %f", countServer, countServerOn, countServerOff, upTimeAvg)
	message := []byte(subject + body)

	//func PlainAuth(identity, username, password, host string) Auth
	auth := smtp.PlainAuth("", from, password, host)

	//func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		panic(err)
	}
	fmt.Println("Check ur mail")
}
