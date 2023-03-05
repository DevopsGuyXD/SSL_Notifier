package main

import (
	"fmt"

	util "github.com/DevopsGuyXD/SSL_Notifier/Utils"
)

func main() {

	util.InitEnvFile()

	// ================ Entry Point ==================
	fmt.Printf("\nWelcome to 'AWS' SSL notifier\n\n")
	util.InitAws()
	GetListOfCertificatesAWS()
}