[![Build Status](https://travis-ci.com/H-Richard/truemail.svg?branch=master)](https://travis-ci.com/H-Richard/truemail)
[![Documentation](https://godoc.org/github.com/H-Richard/truemail?status.svg)](https://godoc.org/github.com/H-Richard/truemail)

# truemail

An email validator package for [Golang](https://golang.org/)

- For regex used, see https://emailregex.com/
- `strict` mode coming soon

## Getting started


### Importing the package

```go
import (
	"github.com/H-Richard/truemail"
)
```

### Validating format

```go
func main() {
  //Validate only format
  v:= truemail.Validator{}
  fmt.PrintLn(v.Validate("r.hong@pm.me"))
}
```

### Validating format and domains

```go
func main() {
  //Validate both format and domain
  v := truemail.Validator{Domains: []string{"pm.me"}}
  fmt.PrintLn(v.Validate("r.hong@pm.me"))
}
```
