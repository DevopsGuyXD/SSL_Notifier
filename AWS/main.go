package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	util "github.com/DevopsGuyXD/SSL_Notifier/Utils"
	"github.com/go-co-op/gocron"
)

var wg sync.WaitGroup

func main(){
	if os.Getenv("IS_CRON") == "true"{
		cron()
	}else{
		singleRun()
	}
}

func cron(){

	s := gocron.NewScheduler(time.Local)

	wg.Add(1)

	s.Cron(os.Getenv("CRON")).Do(func() {
		go startCheck()
	})
	s.StartAsync()

	wg.Wait()
}

func singleRun(){
	startCheck()
}

func startCheck() {
	util.InitEnvFile()

	// ================ Entry Point ==================
	fmt.Printf("\nWelcome to 'AWS' SSL notifier\n\n")
	util.InitAws()
	GetListOfCertificatesAWS()
}
