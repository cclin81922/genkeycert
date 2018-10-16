//    Copyright 2018 cclin
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package genkeycert

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net"
	"os"
	"strings"
	"time"
)

var (
	host      = "dummy"
	validFrom = ""
	validFor  = 365 * 24 * time.Hour
	rsaBits   = 2048
)

// LoadCACertFile ...
func LoadCACertFile() (*x509.Certificate, error) {
	// Read from file
	caCertFile, err := ioutil.ReadFile("pki/ca.cert.pem")

	if err != nil {
		return nil, err
	}

	// PEM decode
	pemBlock, _ := pem.Decode(caCertFile)

	if pemBlock == nil {
		return nil, fmt.Errorf("%s", "pem.Decode failed")
	}

	// Parse x509 Certificate
	caCert, err := x509.ParseCertificate(pemBlock.Bytes)

	if err != nil {
		return nil, err
	}

	return caCert, nil
}

// LoadCAPrivateKeyFile ...
func LoadCAPrivateKeyFile() (*rsa.PrivateKey, error) {
	// Read from file
	caPrivateKeyFile, err := ioutil.ReadFile("pki/ca.key.pem")

	if err != nil {
		return nil, err
	}

	// PEM decode
	pemBlock, _ := pem.Decode(caPrivateKeyFile)

	if pemBlock == nil {
		return nil, fmt.Errorf("%s", "pem.Decode failed")
	}

	// Parse PKCS1 private key
	der := pemBlock.Bytes

	caPrivateKey, err := x509.ParsePKCS1PrivateKey(der)
	if err != nil {
		return nil, err
	}

	return caPrivateKey, nil
}

// MakeClientKey ...
func MakeClientKey() (*rsa.PrivateKey, error) {
	clientPrivateKey, err := rsa.GenerateKey(rand.Reader, rsaBits)

	if err != nil {
		return nil, err
	}

	return clientPrivateKey, nil
}

// MakeClientCertFile ...
// revised version. source from https://golang.org/src/crypto/tls/generate_cert.go
func MakeClientCertFile(caCert *x509.Certificate, caKey, clientKey *rsa.PrivateKey) (*x509.Certificate, []byte) {
	var err error
	var notBefore time.Time
	if len(validFrom) == 0 {
		notBefore = time.Now()
	} else {
		notBefore, err = time.Parse("Jan 2 15:04:05 2006", validFrom)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse creation date: %s\n", err)
			os.Exit(1)
		}
	}

	notAfter := notBefore.Add(validFor)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		log.Fatalf("failed to generate serial number: %s", err)
	}

	// Make a certificate temaplte
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Acme Co"},
		},
		NotBefore: notBefore,
		NotAfter:  notAfter,

		//KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		//ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	hosts := strings.Split(host, ",")
	for _, h := range hosts {
		if ip := net.ParseIP(h); ip != nil {
			template.IPAddresses = append(template.IPAddresses, ip)
		} else {
			template.DNSNames = append(template.DNSNames, h)
		}
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, caCert, publicKey(clientKey), caKey)
	if err != nil {
		log.Fatalf("Failed to create certificate: %s", err)
	}

	clientCert, err := x509.ParseCertificate(derBytes) // NOT SURE ...
	if err != nil {
		panic(err)
	}

	return clientCert, derBytes
}

// SaveClientKeyFile ...
func SaveClientKeyFile(key *rsa.PrivateKey) error {
	keyOut, err := os.OpenFile("dummy.key.pem", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)

	if err != nil {
		return fmt.Errorf("failed to open key.pem for writing: %s", err)
	}

	if err := pem.Encode(keyOut, pemBlockForKey(key)); err != nil {
		return fmt.Errorf("failed to write data to key.pem: %s", err)
	}

	if err := keyOut.Close(); err != nil {
		return fmt.Errorf("error closing key.pem: %s", err)
	}

	return nil
}

// SaveClientCertFile ...
func SaveClientCertFile(cert *x509.Certificate, derBytes []byte) error {
	// TODO remove derBytes

	certOut, err := os.Create("dummy.cert.pem")

	if err != nil {
		return fmt.Errorf("failed to open cert.pem for writing: %s", err)
	}

	if err := pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
		return fmt.Errorf("failed to write data to cert.pem: %s", err)
	}

	if err := certOut.Close(); err != nil {
		return fmt.Errorf("error closing cert.pem: %s", err)
	}

	return nil
}

// source from https://golang.org/src/crypto/tls/generate_cert.go
func publicKey(priv interface{}) interface{} {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &k.PublicKey
	case *ecdsa.PrivateKey:
		return &k.PublicKey
	default:
		return nil
	}
}

// source from https://golang.org/src/crypto/tls/generate_cert.go
func pemBlockForKey(priv interface{}) *pem.Block {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(k)}
	case *ecdsa.PrivateKey:
		b, err := x509.MarshalECPrivateKey(k)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to marshal ECDSA private key: %v", err)
			os.Exit(2)
		}
		return &pem.Block{Type: "EC PRIVATE KEY", Bytes: b}
	default:
		return nil
	}
}
