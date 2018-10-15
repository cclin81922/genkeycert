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
