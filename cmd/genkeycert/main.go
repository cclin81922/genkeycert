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
	"log"

	"github.com/cclin81922/genkeycert/pkg/genkeycert"
)

func main() {
	caCert, err := genkeycert.LoadCACertFile()

	if err != nil {
		log.Fatal(err)
	}

	caKey, err := genkeycert.LoadCAPrivateKeyFile()

	if err != nil {
		log.Fatal(err)
	}

	clientKey, err := genkeycert.MakeClientKey()

	if err != nil {
		log.Fatal(err)
	}

	clientCert, clientCertDerBytes, err := genkeycert.MakeClientCert(caCert, caKey, clientKey)

	if err != nil {
		log.Fatal(err)
	}

	if err := genkeycert.SaveClientKeyFile(clientKey); err != nil {
		log.Fatal(err)
	}

	if err := genkeycert.SaveClientCertFile(clientCert, clientCertDerBytes); err != nil {
		log.Fatal(err)
	}

}
