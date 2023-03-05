package main

import (
	"fmt"

	util "github.com/DevopsGuyXD/SSL_Notifier/Utils"
)

func main() {

	util.InitEnvFile()

	// ================== Entry Point ====================
	fmt.Printf("\n\nWelcome to 'Azure' SSL notifier\n\n")
	util.InitAzure()
	GetListOfCertificatesAzure()
}