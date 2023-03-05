package util

import (
	"log"
	"os"
	"os/exec"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/joho/godotenv"
)

func CheckForMajorErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func InitEnvFile(){
	err := godotenv.Load(".env"); CheckForMajorErr(err)
}

func InitAws(){
	InitEnvFile()

	exec.Command("aws", "configure", "set", "aws_access_key_id", os.Getenv("ACCESS_KEY_ID")).Output()
	exec.Command("aws", "configure", "set", "aws_secret_access_key", os.Getenv("ACCESS_KEY_SECRET")).Output()
	exec.Command("aws", "configure", "set", "region", os.Getenv("AWS_REGION")).Output()

	_, err := session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config: aws.Config{
			Region: aws.String(os.Getenv("AWS_REGION")),
		},
	})
	CheckForMajorErr(err)
}

func InitAzure(){
	InitEnvFile()

	_, err := exec.Command("az", "login", "--service-principal", "-u", os.Getenv("CLIENT_ID"), "-p", os.Getenv("SECRET_VALUE"), "--tenant", os.Getenv("TENANT_ID")).Output(); if err != nil{
		log.Println(err)
	}
}