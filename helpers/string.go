package helpers

import "strings"

type String struct {
	String string
}

func (s *String) IsError() bool {
	return s.String != "" || s.String == "ERROR : "
}

func (s *String) Append(str string) *String {
	(*s).String += str

	return s
}

func (s *String) AppendIf(condition bool, ifTrue string, ifFalse string) *String {
	if condition {
		(*s).String += ifTrue
	} else {
		(*s).String += ifFalse
	}

	return s
}

func (s *String) IsEmpty() bool {
	return strings.Trim(s.String, " ") == ""
}
