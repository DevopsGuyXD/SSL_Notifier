package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	util "github.com/DevopsGuyXD/SSL_Notifier/Utils"
	"github.com/go-co-op/gocron"
)

// ================ Wait group ================
var wg sync.WaitGroup

// =================== MAIN ===================
func main(){

	if os.Getenv("IS_CRON") == "true"{
		cron()
	}else{
		singleRun()
	}
}

// =================== CRON ===================
func cron(){

	s := gocron.NewScheduler(time.Local)

	wg.Add(1)

	s.Cron(os.Getenv("CRON")).Do(func() {
		go startCheck()
	})
	s.StartAsync()

	wg.Wait()
}

// ============= Single execution =============
func singleRun(){
	startCheck()
}

// =============== Start checker ==============
func startCheck() {
	util.InitEnvFile()

	// ------------ Entry Point ---------------
	fmt.Printf("\nWelcome to 'AWS' SSL notifier\n\n")
	util.InitAws()
	GetListOfCertificatesAWS()
}
