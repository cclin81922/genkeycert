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
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
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

func main() {
	// TODO
}
