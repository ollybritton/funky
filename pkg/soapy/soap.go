/*
Soapy is a package that allows the creation of liquidsoap files and code through the Go syntax.
*/

package soapy

import "strings"

// CanSoap means that the supplied type can output liquidsoap code.
type CanSoap interface {
	Out() string
}

// Soap is the main type that allows an interface between Go and liquidsoap.
type Soap struct {
	rawSoap []string
}

// Add will add a string to the raw liquidsoap contained within the Soap struct.
func (s *Soap) Add(code string) {
	s.rawSoap = append(s.rawSoap, code)
}

// Out will get all the liquidsoap code associated with the Soap and return it as a string.
func (s Soap) Out() string {
	return strings.Join(s.rawSoap, "\n")
}

// NewSoap returns a new soap struct ready for usage.
func NewSoap() *Soap {
	s := new(Soap)
	s.rawSoap = []string{}

	return s
}
