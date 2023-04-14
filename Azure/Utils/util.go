package util

import (
	"log"
	"os"
	"os/exec"

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

func InitAzure(){
	InitEnvFile()

	_, err := exec.Command("az", "login", "--service-principal", "-u", os.Getenv("CLIENT_ID"), "-p", os.Getenv("SECRET_VALUE"), "--tenant", os.Getenv("TENANT_ID")).Output(); if err != nil{
		log.Println(err)
	}
}