package main

import (
	"crypto/tls"
	"fmt"
	"math"
	"os"
	"os/exec"
	"strings"
	"time"

	util "github.com/DevopsGuyXD/SSL_Notifier/Utils"
	gomail "gopkg.in/mail.v2"
)

// ========================== Calculate days to expire ==============================
func DaysToExpireAWS(i int, expiry_date []float64) int{
	current_date := time.Now()

	parse_date := int64(expiry_date[i])
	parse_date_human_readable := time.Unix(parse_date, 0)
	parse_date_string := parse_date_human_readable.String()
	date, err := time.Parse("2006-01-02", parse_date_string[0:10]); util.CheckForMajorErr(err)

	difference := date.Sub(current_date)

	days_till_expiry := int(difference.Hours()/24)

	return days_till_expiry
}

// ========= Calculate days to expire and send notification control center ============
func GetDaysLeftForExpiryAWS(cert_id []string, cert_domain_name []string, expiry_date []float64) {

	for i := 0; i < len(expiry_date); i++ {

		days_till_expiry := DaysToExpireAWS(i, expiry_date)

		if !math.Signbit(float64(days_till_expiry)) && days_till_expiry < 15 {
			SendNotificationAWS(cert_id[i], cert_domain_name[i], days_till_expiry)
		}
	}

	fmt.Println("\nCompleted successfully")
}

// ================================= Send notification =================================
func SendNotificationAWS(cert_arn string, cert_domain_name string, days_till_expiry int){

	util.InitEnvFile()

	account_id, err := exec.Command("aws", "sts", "get-caller-identity", "--query", "Account", "--output", "text").Output(); util.CheckForMajorErr(err)
	
	email := gomail.NewMessage()

	recipients := os.Getenv("RECEIPIENTS")
	recipients_parsed := strings.Split(recipients,",")

	for i := 0; i < len(recipients_parsed); i++{
		email_subject := "SSL renewal reminder - AWS"
		email_body := fmt.Sprintf("<h4>%v</h4>%v<br><br><b>Account_ID:</b> %v<br><b>Certificate_ARN:</b> %v<br><b>Certificate_Domain:</b> %v<br><h1 style=`text-align:center;font-size:80px;color:#FF9900;`>%v<div style=`font-size:20px;color:black;`>Days to expire</div></h1>","Greetings user,","The below certificate is due for renewal. Please take the necessary action at the earliest.",string(account_id),cert_arn,cert_domain_name,days_till_expiry)
		email_connection := os.Getenv("EMAIL_CONNECTION")
		port := 587

		email.SetHeader("From", os.Getenv("EMAIL_SENDER_ID"))
		email.SetHeader("To", recipients_parsed[i])
		email.SetHeader("Subject", email_subject)
		email.SetBody("text/html", email_body)

		d := gomail.NewDialer(email_connection, port, os.Getenv("EMAIL_SENDER_ID"), os.Getenv("EMAIL_SENDER_PASSWORD"))

		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

		err = d.DialAndSend(email); util.CheckForMajorErr(err)

		fmt.Printf("%v: Notification sent\n", cert_arn)
	}
}