package soapy

import "fmt"

// Output defines a type that can be used to generate the code that outputs the stream in liquidsoap.
type Output struct {
	Name  string
	Desc  string
	Genre string

	Host     string
	Port     int
	Password string

	SampleRate int
	Bitrate    int
	Public     bool
	Stereo     bool
	ID3v2      bool
	Encoding   string
	Mount      string
	ID         string
	Protocol   string
}

// Out returns the liquidsoap code associated with the output.
func (o *Output) Out() string {
	return fmt.Sprintf(
		`output.icecast(%%mp3(samplerate=%d, stereo=%t, bitrate=%d, id3v2=%t), id="%v", host = "%v", port = %d, password = "%v", mount = "%v", name = "%v", description = "%v", genre = "%v", public = %t, encoding = "%v", protocol="%v", radio)`,
		o.SampleRate,
		o.Stereo,
		o.Bitrate,
		o.ID3v2,
		o.ID,
		o.Host,
		o.Port,
		o.Password,
		o.Mount,
		o.Name,
		o.Desc,
		o.Genre,
		o.Public,
		o.Encoding,
		o.Protocol,
	)
}

// NewOutput will return a Output struct with some opinionated defaults.
func NewOutput(name string, desc string, genre string, host string, port int, password string) Output {
	o := new(Output)
	o.Name, o.Desc, o.Genre, o.Host, o.Port, o.Password = name, desc, genre, host, port, password

	o.SampleRate = 44100
	o.Bitrate = 128
	o.Public = false
	o.Stereo = true
	o.ID3v2 = true
	o.Encoding = "UTF-8"
	o.Mount = "/radio.mp3"
	o.ID = GetID(name)
	o.Protocol = "icy"

	return *o

}
