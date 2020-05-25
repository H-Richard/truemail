package truemail_test

import (
	"testing"

	"github.com/H-Richard/truemail"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		email    string
		expected error
	}{
		{email: "", expected: truemail.ErrFor},
		{email: "@", expected: truemail.ErrFor},
	}
	v := truemail.Validator{}
	for _, c := range tests {
		if actual := v.Validate(c.email); actual != c.expected {
			t.Errorf("Test failed for %s\nexpected: %s\nactual: %s", c.email, c.expected, actual)
		}
	}
}

func TestInvalidDomainFormat(t *testing.T) {
	tests := []struct {
		email    string
		expected error
	}{
		{email: "richard@aaaa", expected: truemail.ErrFor},
		{email: "richard@bbbbb,", expected: truemail.ErrFor},
		{email: "richard@bbbbb.,", expected: truemail.ErrFor},
	}
	v := truemail.Validator{}
	for _, c := range tests {
		if actual := v.Validate(c.email); actual != c.expected {
			t.Errorf("Test failed for %s\nexpected: %s\nactual: %s", c.email, c.expected, actual)
		}
	}
}
