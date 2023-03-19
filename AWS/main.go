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

	// is_cron, err := strconv.ParseBool(strings.ToLower(os.Getenv("IS_CRON"))); util.CheckForMajorErr(err)

	if os.Getenv("IS_CRON") == "true"{
		cron()
	}else{
		singleRun()
	}
}

// ============= Single execution =============
func singleRun(){
	startCheck()
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

// =============== Start checker ==============
func startCheck() {
	util.InitEnvFile()

	// ------------ Entry Point ---------------
	fmt.Printf("\nWelcome to 'AWS' SSL notifier\n\n")
	util.InitAws()
	GetListOfCertificatesAWS()
}
