package services

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"
	"time"
)

type X509Service struct {
	configService *ConfigService
}

func NewX509Service(configService *ConfigService) *X509Service {
	return &X509Service{
		configService: configService,
	}
}

type CertificateInfo struct {
	CommonName     string   `json:"common_name"`
	Organization   string   `json:"organization"`
	Email          string   `json:"email"`
	SerialNumber   string   `json:"serial_number"`
	Issuer         string   `json:"issuer"`
	PSD2Roles      []string `json:"psd2_roles,omitempty"`
	IsValid        bool     `json:"is_valid"`
	ValidationError string   `json:"validation_error,omitempty"`
}

func (x *X509Service) ValidateCertificate(encodedCert string) (*CertificateInfo, error) {
	cert, err := x.parseCertificate(encodedCert)
	if err != nil {
		return &CertificateInfo{
			IsValid:         false,
			ValidationError: err.Error(),
		}, err
	}

	validationErr := x.checkCertificateValidity(cert)
	
	var email string
	if len(cert.EmailAddresses) > 0 {
		email = cert.EmailAddresses[0]
	}

	var organization string
	if len(cert.Subject.Organization) > 0 {
		organization = cert.Subject.Organization[0]
	}

	psd2Roles := x.extractPSD2Roles(cert)

	info := &CertificateInfo{
		CommonName:   cert.Subject.CommonName,
		Organization: organization,
		Email:        email,
		SerialNumber: cert.SerialNumber.String(),
		Issuer:       cert.Issuer.CommonName,
		PSD2Roles:    psd2Roles,
		IsValid:      validationErr == nil,
	}

	if validationErr != nil {
		info.ValidationError = validationErr.Error()
	}

	return info, nil
}

func (x *X509Service) parseCertificate(encodedCert string) (*x509.Certificate, error) {
	var certBytes []byte
	
	if strings.Contains(encodedCert, "-----BEGIN CERTIFICATE-----") {
		block, _ := pem.Decode([]byte(encodedCert))
		if block == nil {
			return nil, errors.New("failed to parse PEM certificate")
		}
		certBytes = block.Bytes
	} else {
		return nil, errors.New("DER format not yet supported")
	}

	cert, err := x509.ParseCertificate(certBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse certificate: %w", err)
	}

	return cert, nil
}

func (x *X509Service) checkCertificateValidity(cert *x509.Certificate) error {
	now := time.Now()
	if now.Before(cert.NotBefore) {
		return errors.New("certificate not yet valid")
	}
	if now.After(cert.NotAfter) {
		return errors.New("certificate has expired")
	}
	
	return nil
}

func (x *X509Service) extractPSD2Roles(cert *x509.Certificate) []string {
	var roles []string
	
	
	for _, ext := range cert.Extensions {
		if ext.Id.String() == "0.4.0.19495.2" {
			roles = append(roles, "PSP_AS", "PSP_PI", "PSP_AI", "PSP_IC")
			break
		}
	}
	
	return roles
}

func (x *X509Service) GetRSAPublicKey(encodedCert string) (*rsa.PublicKey, error) {
	cert, err := x.parseCertificate(encodedCert)
	if err != nil {
		return nil, err
	}

	rsaKey, ok := cert.PublicKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("certificate does not contain RSA public key")
	}

	return rsaKey, nil
}

func (x *X509Service) ExtractSubjectInfo(encodedCert string) (map[string]string, error) {
	cert, err := x.parseCertificate(encodedCert)
	if err != nil {
		return nil, err
	}

	info := make(map[string]string)
	info["common_name"] = cert.Subject.CommonName
	info["country"] = strings.Join(cert.Subject.Country, ",")
	info["organization"] = strings.Join(cert.Subject.Organization, ",")
	info["organizational_unit"] = strings.Join(cert.Subject.OrganizationalUnit, ",")
	info["locality"] = strings.Join(cert.Subject.Locality, ",")
	info["province"] = strings.Join(cert.Subject.Province, ",")
	
	if len(cert.EmailAddresses) > 0 {
		info["email"] = cert.EmailAddresses[0]
	}

	return info, nil
}
