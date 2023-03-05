package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	util "github.com/DevopsGuyXD/SSL_Notifier/Utils"
	"github.com/gorilla/mux"
	"github.com/go-co-op/gocron"
)

var wg sync.WaitGroup

func main() {
	util.InitEnvFile()

	// ================ Entry Point ==================
	fmt.Printf("\nWelcome to 'AWS' SSL notifier\n\n")
	util.InitAws()
	GetListOfCertificatesAWS()
}

func healthCheck(){
	router := mux.NewRouter();
	router.HandleFunc("/", Health).Methods("GET")
	http.ListenAndServe(":80",router)
}

func Health(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	fmt.Fprintf(w,"%v = %v\n\nExecutes everyday at 8:00 AM IST","Healty",time.Now())
}

func init(){

	s := gocron.NewScheduler(time.Local)

	wg.Add(2)

	s.Cron("30 02 * * *").Do(func() {
		go main()
	})
	s.StartAsync()
	go healthCheck()

	wg.Wait()
}