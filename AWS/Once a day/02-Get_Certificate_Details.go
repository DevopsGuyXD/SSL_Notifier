package main

import (
	"encoding/json"
	"fmt"
	"os/exec"

	util "github.com/DevopsGuyXD/SSL_Notifier/Utils"
)

// ====================================== STRUCTS ========================================
type Certificate_Details struct{
	Certificate *Certificate `json:"certificate,omitempty"`
}

type Certificate struct{
	CertificateArn string `json:"certificatearn,omitempty"`
	DomainName string `json:"domainname,omitempty"`
	SubjectAlternativeNames []string `json:"subjectalternativenames,omitempty"`
	DomainValidationOptions []*Domain_validation_options `json:"domainvalidationoptions,omitempty"`
	Subject string `json:"subject,omitempty"`
	Issuer string `json:"issuer,omitempty"`
	CreatedAt string `json:"createdat,omitempty"`
	IssuedAt string `json:"issuedat,omitempty"`
	Status string `json:"status,omitempty"`
	NotBefore string `json:"notbefore,omitempty"`
	NotAfter float64 `json:"notafter,omitempty"`
	KeyAlgorithm string `json:"keyalgorithm,omitempty"`
	SignatureAlgorithm string `json:"signaturealgorithm,omitempty"`
	InUseBy []string `json:"inuseby,omitempty"`
	Type string `json:"type,omitempty"`
	KeyUsages []string `json:"keyusages,omitempty"`
	ExtendedKeyUsages []string `json:"extendedkeyusages,omitempty"`
	RenewalEligibility string `json:"renewaleligibility,omitempty"`
	Options *Certificate_transparency_logging_preference `json:"options,omitempty"`
}

type Domain_validation_options struct{
	DomainName string `json:"domainname,omitempty"`
	ValidationDomain string `json:"validationdomain,omitempty"`
	ValidationStatus string `json:"validationstatus,omitempty"`
	ResourceRecord *Resource_record `json:"ResourceRecord,omitempty"`
	ValidationMethod string `json:"validationmethod,omitempty"`
}

type Resource_record struct{
	Name string `json:"Name,omitempty"`
	Type string `json:"Type,omitempty"`
	Value string `json:"Value,omitempty"`
}

type Certificate_transparency_logging_preference struct{
	CertificateTransparencyLoggingPreference string `json:"certificatetransparencyloggingpreference,omitempty"`
}


// ================================= Get certificate detials ===================================
func GetCertifcateDetailsAWS(certificate *All_Certificates) {

	var certificate_details *Certificate_Details
	var cert_arn []string
	var cert_domain_name []string
	var expiry_date []float64

	for i := 0; i < certificate.TotalNummberOfCertificates(); i++{
		res, err := exec.Command("aws","acm","describe-certificate","--certificate-arn",certificate.CertificateSummaryList[i].CertificateArn).Output(); util.CheckForMajorErr(err)
		json.Unmarshal(res, &certificate_details)

		fmt.Printf("(%v/%v)",i+1, certificate.TotalNummberOfCertificates())
		fmt.Println(certificate_details.Certificate.CertificateArn)

		if certificate_details.Certificate.Status == "ISSUED"{
			cert_arn = append(cert_arn, certificate_details.Certificate.CertificateArn)
			cert_domain_name = append(cert_domain_name, certificate_details.Certificate.DomainName)
			expiry_date = append(expiry_date,certificate_details.Certificate.NotAfter)
		}
	}

	fmt.Println("")
	GetDaysLeftForExpiryAWS(cert_arn, cert_domain_name, expiry_date)
}