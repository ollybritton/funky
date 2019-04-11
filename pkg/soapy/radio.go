package soapy

import "strings"

// Radio represents the object that actually gets streamed by liquidsoap
type Radio struct {
	raw []string
}

// Add adds raw liquidsoap code to the radio struct
func (r *Radio) Add(code string) {
	r.raw = append(r.raw, code)
}

// Out exports the Radio struct into liquidsoap code
func (r *Radio) Out() string {
	return strings.Join(r.raw, "\n")
}

// AddCategories adds a set of categories to the soap
func (s *Soap) AddCategories(playlists []Playlist) {
	r := Radio{}

	code := "radio = rotate(["

	for i, playlist := range playlists {
		if i == len(playlists)-1 {
			code += playlist.ID
		} else {
			code += playlist.ID + ","
		}
	}

	code += "])"

	r.Add(code)
	s.Add(r.Out())
}
