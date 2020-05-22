package truemail

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	formatRegexp, _ = regexp.Compile("(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$)")

	//ErrFor format error
	ErrFor = errors.New("invalid format")
	//ErrDom domain error
	ErrDom = errors.New("invalid domain")
)

//Validator struct
type Validator struct {
	Domains []string
	Strict  bool
}

//Validate : validates emails
func (v *Validator) Validate(email string) error {
	if !formatRegexp.MatchString(email) {
		return ErrFor
	}
	if domains := v.Domains; len(domains) > 0 {
		domainRegexp, _ := regexp.Compile(fmt.Sprintf(".*(%s)$", strings.Join(domains, "|")))
		if !domainRegexp.MatchString(email) {
			return ErrDom
		}
	}
	return nil
}
