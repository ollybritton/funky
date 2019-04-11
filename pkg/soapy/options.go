package soapy

import (
	"fmt"

	"github.com/spf13/cast"
)

// SetOptionBool will set a boolean option in liquidsoap.
// Example: set("log.stdout", true)
func (s *Soap) SetOptionBool(optionName string, optionValue bool) {
	code := fmt.Sprintf(`set("%v",%v)`, optionName, cast.ToString(optionValue))
	s.Add(code)
}

// SetOptionString will a set a string option in liquidsoap.
// Example: set("server.telnet.bind_addr","127.0.0.1")
func (s *Soap) SetOptionString(optionName string, optionValue string) {
	code := fmt.Sprintf(`set("%v",%v)`, optionName, optionValue)
	s.Add(code)
}

// SetOptionInt will set an integer option in liquidsoap.
// Example: set("server.telnet.port", 8004)
func (s *Soap) SetOptionInt(optionName string, optionValue int) {
	code := fmt.Sprintf(`set("%v",%v)`, optionName, cast.ToString(optionValue))
	s.Add(code)
}

// SetOptionStringArray will set a option in liquidsoap of the form
//  set("option", [string, string2, string3])
func (s *Soap) SetOptionStringArray(optionName string, optionValue []string) {
	stringArray := "["

	for i, val := range optionValue {
		if i != len(optionValue)-1 {
			stringArray += fmt.Sprintf(`"%v",`, val)
		} else {
			stringArray += fmt.Sprintf(`"%v"`, val)
		}
	}

	stringArray += "]"

	code := fmt.Sprintf(`set("%v",%v)`, optionName, stringArray)
	s.Add(code)
}
