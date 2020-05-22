package truemail

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	formatRegexp, _ = regexp.Compile("(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$)")
)

//Validator struct
type Validator struct {
	Domains []string
	Strict  bool
}

//Validate : validates emails
func (v *Validator) Validate(email string) error {
	if !formatRegexp.MatchString(email) {
		return errors.New("invalid format")
	}
	if domains := v.Domains; len(domains) > 0 {
		domainRegexp, _ := regexp.Compile(fmt.Sprintf(".*(%s)$", strings.Join(domains, "|")))
		if !domainRegexp.MatchString(email) {
			return errors.New("invalid domain")
		}

	}
	return nil
}
