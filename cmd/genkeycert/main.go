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

package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
)

func loadCACertFile() *x509.Certificate {
	caCertFile, err := ioutil.ReadFile("pki/ca.cert.pem")

	if err != nil {
		panic(err)
	}

	pemBlock, _ := pem.Decode(caCertFile)

	if pemBlock == nil {
		panic("pem.Decode failed")
	}

	caCert, err := x509.ParseCertificate(pemBlock.Bytes)

	if err != nil {
		panic(err)
	}

	return caCert
}

func loadCAPrivateKeyFile() *rsa.PrivateKey {
	caPrivateKeyFile, err := ioutil.ReadFile("pki/ca.key.pem")

	if err != nil {
		panic(err)
	}

	pemBlock, _ := pem.Decode(caPrivateKeyFile)

	if pemBlock == nil {
		panic("pem.Decode failed")
	}

	der, err := x509.DecryptPEMBlock(pemBlock, []byte(""))
	if err != nil {
		panic(err)
	}

	caPrivateKey, err := x509.ParsePKCS1PrivateKey(der)
	if err != nil {
		panic(err)
	}

	return caPrivateKey
}

func makeClientKeyFile() *rsa.PrivateKey {
	clientPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		panic(err)
	}

	return clientPrivateKey
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

func main() {
	// TODO
}
