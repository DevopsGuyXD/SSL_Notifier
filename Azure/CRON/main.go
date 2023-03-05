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

func main() {

	util.InitEnvFile()

	// ================== Entry Point ====================
	fmt.Printf("\n\nWelcome to 'Azure' SSL notifier\n\n")
	util.InitAzure()
	GetListOfCertificatesAzure()
}

func init(){

	util.InitEnvFile()

	s := gocron.NewScheduler(time.Local)

	wg.Add(1)

	s.Cron(os.Getenv("CRON")).Do(func() {
		go main()
	})
	s.StartAsync()

	wg.Wait()
}
