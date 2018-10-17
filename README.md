# Installation

```
go get -u github.com/cclin81922/genkeycert/cmd/genkeycert
export PATH=$PATH:~/go/bin
```

# Command Line Usage

```
genkeycert
```

Verify the output

```
# inspect certificate
openssl x509 -in dummy.cert.pem -text -noout

# test osbapibaas
osbapibaas -port=8443
curl -k --key dummy.key.pem --cert dummy.cert.pem https://localhost.localdomain:8443/echo -d "hi"
```

# Package Usage

```
import github.com/cclin81922/genkeycert/pkg/genkeycert

func demo(cn string) {
    keyStr, certStr, err := genkeycert.MakeClientKeyCert(cn)
}
```

# Related Resources

* [Signing certificate request with certificate authority in golang](https://stackoverflow.com/questions/42643048/signing-certificate-request-with-certificate-authority)
* [Adding an extension (subjectAlName) to a x509 certificate in golang](https://stackoverflow.com/questions/26441547/go-how-do-i-add-an-extension-subjectaltname-to-a-x509-certificate)
* [Certificate and key formats](http://jianiau.blogspot.com/2015/07/openssl-key-and-certificate-conversion.html)
* [Error - x509: no DEK-Info header in block](https://stackoverflow.com/questions/32981821/no-dek-info-header-in-block-when-attempting-to-read-encrypted-private-key)
* [Error - x509: certificate specifies an incompatible key usage](https://github.com/hashicorp/vault/issues/846)