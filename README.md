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
osbapibaas -port=8443
curl -k --key dummy.key.pem --cert dummy.cert.pem https://localhost.localdomain:8443/echo -d "hi"
```

# Package Usage

```
import github.com/cclin81922/genkeycert/pkg/genkeycert

func demo(cn string) {
    // TODO
}
```

# Related Resources

* [Signing certificate request with certificate authority in golang](https://stackoverflow.com/questions/42643048/signing-certificate-request-with-certificate-authority)
* [Error - x509: no DEK-Info header in block](https://stackoverflow.com/questions/32981821/no-dek-info-header-in-block-when-attempting-to-read-encrypted-private-key)
* [Certificate and key formats](http://jianiau.blogspot.com/2015/07/openssl-key-and-certificate-conversion.html)