package soapy

import "strings"

// GetID returns an appropriate ID for the playlist name given.
func GetID(name string) string {
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "_")

	return name
}
