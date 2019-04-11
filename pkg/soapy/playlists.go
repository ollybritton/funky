package soapy

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"strings"
)

var (
	AllowedFormats = []string{".mp3", ".wav"}
)

// Playlist represents a set of songs that can be played.
type Playlist struct {
	Name    string   // Human readable name for the playlist, e.g. "Modern Music"
	ID      string   // The name of the variable that the playlist is assigned to when generated.
	Path    string   // The path to the media files
	Reload  string   // The input to reload_mode="" in liquidsoap
	Mode    string   // The input to mode="" in liquidsoap
	Effects []string // What functions to wrap the playlist in, such as "audio_to_stereo"
}

// GeneratePlaylistFile generates a file containing the URI of each song.
func (p *Playlist) GeneratePlaylistFile() string {
	files, err := ioutil.ReadDir(p.Path)
	if err != nil {
		log.Fatal(err)
	}

	uris := []string{}
	for _, file := range files {
		fileName := file.Name()

		for _, val := range AllowedFormats {
			if filepath.Ext(fileName) == val {
				uris = append(uris, fileName)
			}
		}

	}

	data := []byte(strings.Join(uris, "\n"))
	playlistFilePath := path.Join(p.Path, "/_meta.list")

	err = ioutil.WriteFile(playlistFilePath, data, 0666)
	if err != nil {
		log.Fatal(err)
	}

	return playlistFilePath

}

// Out returns a string representation of the playlist as liquidsoap code.
func (p *Playlist) Out() string {
	var code string

	playlistFilePath := p.GeneratePlaylistFile()

	code = fmt.Sprintf(`playlist(reload_mode="%v",mode="%v","%v")`, p.Reload, p.Mode, playlistFilePath)

	for _, val := range p.Effects {
		code = val + "(" + code + ")"
	}

	code = p.ID + " = " + code

	return code
}

// NewPlaylist returns a new playlist struct with some better initial values
func NewPlaylist(name string, path string, make_safe bool) Playlist {
	p := new(Playlist)

	p.Name = name
	p.Path = path

	p.ID = "__playlist__" + GetID(name)
	p.Reload = "watch"
	p.Mode = "randomize"

	p.Effects = []string{"audio_to_stereo"}

	if make_safe {
		p.Effects = append(p.Effects, "mksafe")
	}

	return *p

}

// AddPlaylist will add a Playlist to a Soap.
func (s *Soap) AddPlaylist(p Playlist) {
	s.Add(p.Out())
}
