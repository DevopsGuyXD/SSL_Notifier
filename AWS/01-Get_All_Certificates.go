package main

import (
	"encoding/json"
	"fmt"
	"os/exec"

	util "github.com/DevopsGuyXD/SSL_Notifier/Utils"
)

// ===================================== AWS Structs ======================================
type All_Certificates struct {
	CertificateSummaryList []*certificateSummaryList `json:"certificatesummarylist,omitempty"`
}

type certificateSummaryList struct {
	CertificateArn string `json:"certificatearn,omitempty"`
	DomainName     string `json:"domainName,omitempty"`
}


// ================================= AWS Structs Methods ==================================
func (ac All_Certificates) TotalNummberOfCertificates() int{

	total_num_certs := len(ac.CertificateSummaryList)
	return total_num_certs
}


// ================================ List AWS certificates =================================
func GetListOfCertificatesAWS() {

	var certificate *All_Certificates

	fmt.Printf("  ◦ Checking...\n\n")
	list_all_certs, err := exec.Command("aws", "acm", "list-certificates").Output(); util.CheckForMajorErr(err)
	err = json.Unmarshal(list_all_certs, &certificate); util.CheckForMajorErr(err)

	GetCertifcateDetailsAWS(certificate)
}