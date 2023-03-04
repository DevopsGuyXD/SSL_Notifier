package main

import (
	"encoding/json"
	"fmt"
	"os/exec"

	util "github.com/DevopsGuyXD/SSL_Notifier/Utils"
)

// =========================================== Azure Structs =====================================
type KeyVaulltDetails struct{
	Id string `json:"id,omitempty"`
	Location string `json:"location,omitempty"`
	Name string `json:"name,omitempty"`
	ResourceGroup string `json:"resourcegroup,omitempty"`
}

type CertificateDetails struct{
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	X509ThumbprintHex string `json:"x509thumbprinthex,omitempty"`
	Attributes *Attributes `json:"attributes,omitempty"`
}

type Attributes struct{
	Created string `json:"created,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Expires string `json:"expires,omitempty"`
	NotBefore string `json:"notbefore,omitempty"`
} 

// ===================================== List Azure certificates ==================================
func GetListOfCertificatesAzure(){

	var keyVaults []*KeyVaulltDetails
	var certificateDetails []*CertificateDetails

	fmt.Printf("Checking...\n\n")
	kvs, err := exec.Command("az","keyvault","list").Output(); util.CheckForMajorErr(err)
	err = json.Unmarshal(kvs, &keyVaults); util.CheckForMajorErr(err)

	for i := 0; i < len(keyVaults); i++{
		kv_details, err := exec.Command("az","keyvault","certificate","list","--vault-name",keyVaults[i].Name).Output(); util.CheckForMajorErr(err)
		err = json.Unmarshal(kv_details, &certificateDetails); util.CheckForMajorErr(err)

		fmt.Println(keyVaults[i].Name)

		if len(kv_details) > 4{
			GetDaysLeftForExpiryAzure(certificateDetails)
		}
	}

	fmt.Println("")
	fmt.Println("Completed successfully")
}