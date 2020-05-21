package truemail

import (
	"errors"
	"regexp"
)

var (
	formatRegexp, _ = regexp.Compile("(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$)")
)

func Validate(email string) error {
	if !formatRegexp.MatchString(email) {
		return errors.New("invalid format")
	}
	return nil
}
