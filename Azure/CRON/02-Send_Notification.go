package main

import (
	"crypto/tls"
	"fmt"
	"math"
	"net/url"
	"os"
	"strings"
	"time"

	util "github.com/DevopsGuyXD/SSL_Notifier/Utils"
	gomail "gopkg.in/mail.v2"
)

// ========================== Calculate days to expire ==============================
func DaysToExpireAzure(i int, expiry_date string) int{
	current_date := time.Now()

	date, err := time.Parse("2006-01-02", expiry_date[0:10]); util.CheckForMajorErr(err)
	difference := date.Sub(current_date)
	days_till_expiry := int(difference.Hours()/24)

	return days_till_expiry
}

// ========== Calculate days to expire and send notification control center =========
func GetDaysLeftForExpiryAzure(certificateDetails []*CertificateDetails){
	for i := 0; i < len(certificateDetails); i++{
		days_till_expiry := DaysToExpireAzure(i, certificateDetails[i].Attributes.Expires)

		parse_id, err := url.Parse(certificateDetails[i].Id); util.CheckForMajorErr(err)
		replacer := strings.NewReplacer(".vault.azure.net","")
		remove_url_from_kv := replacer.Replace(parse_id.Host)

		if !math.Signbit(float64(days_till_expiry)) && days_till_expiry < 15{
			SendNotificationAzure( certificateDetails[i].Id, certificateDetails[i].Name, remove_url_from_kv, days_till_expiry)
		}
	}
}

// ================================= Send notification ===============================
func SendNotificationAzure(certId string, certName string, key_vault_name string, days_till_expiry int){

	util.InitEnvFile()
	
	email := gomail.NewMessage()

	email_subject := "SSL renewal reminder - Azure"
	email_body := fmt.Sprintf("<h4>%v</h4>%v<br><br><b>Key_Vault_Name:</b> %v<br><b>Certificate_Id:</b> %v<br><b>Certificate_Name:</b> %v<br><h1 style=`text-align:center;font-size:80px;color:#0080FF;`>%v<div style=`font-size:20px;color:black;`>Days to expire</div></h1>","Greetings user,","The below certificate is due for renewal. Please take the necessary action at the earliest.",key_vault_name,string(certId),certName,days_till_expiry)
	email_connection := os.Getenv("EMAIL_CONNECTION")
	port := 587

	email.SetHeader("From", os.Getenv("EMAIL_SENDER_ID"))
	email.SetHeader("To", os.Getenv("RECEIPIENT_MAIN"))
	email.SetHeader("Cc", os.Getenv("RECEIPIENT_CC_1"), os.Getenv("RECEIPIENT_CC_2"))
	email.SetHeader("Subject", email_subject)
	email.SetBody("text/html", email_body)

	d := gomail.NewDialer(email_connection, port, os.Getenv("EMAIL_SENDER_ID"), os.Getenv("EMAIL_SENDER_PASSWORD"))

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err := d.DialAndSend(email); util.CheckForMajorErr(err)

	fmt.Printf("%v: Notification sent\n", certId)
}