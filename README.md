# Installation

```
go get -u github.com/cclin81922/genkeycert/cmd/genkeycert
export PATH=$PATH:~/go/bin
```

# Command Line Usage

```
genkeycert -cn=dummy
```

# Package Usage

```
import github.com/cclin81922/genkeycert/pkg/genkeycert

func demo(cn string) {
    // TODO
}
```

# Related Resources

* [Signing certificate request with certificate authority](https://stackoverflow.com/questions/42643048/signing-certificate-request-with-certificate-authority)
* [x509: no DEK-Info header in block](https://stackoverflow.com/questions/32981821/no-dek-info-header-in-block-when-attempting-to-read-encrypted-private-key)