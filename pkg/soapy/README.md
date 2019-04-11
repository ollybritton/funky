# Soapy
Soapy is a package that allows you to generate Liquidsoap code programmatically through a Go API.

## Usage
### Soaps & Playlists
To start, you define a new `Soap`, a struct holding all the information about the Liquidsoap program and methods associated with it.

```go
soap := soapy.NewSoap()
```

You can then use methods such as

```go
soap.SetOptionBool("log.stdout", true)
// Writes the 'set("log.stdout", true)' option in Liquidsoap.
```

Playlists also work in a similar way. You first define a new `Playlist`, and then can add it to the soap.

```go
playlist := soapy.NewPlaylist(name, path, safe)
// name : Name of the playlist
// path : Path to the files associated with the playlist
// safe : Whethere to mark the playlist as unfalliable

soap.AddPlaylist(playlist)
```

For example, let's see what the following code does:

```go
// Create a new soap
soap := soapy.NewSoap()

// Set some options
soap.SetOptionBool("log.stdout", true)
soap.SetOptionBool("log.file", false)

// Define a new playlist and add it to the soap
// Note: This will only work if you have the directory 'test/media/music'
playlist := soapy.NewPlaylist("Music", "test/media/music", true)
soap.AddPlaylist(playlist)

// Print the Soap as Liquidsoap code.
// .Out() works on many of the different Soapy types.
fmt.Println(soap.Out())
```

It generates the following

```
set("log.stdout",true)
set("log.file",false)
__playlist__music = mksafe(audio_to_stereo(playlist(reload_mode="watch",mode="randomize","test/media/music/_meta.list")))
```

If you were to run the code, it wouldn't do much. That's because it isn't actually doing anything with the playlist that has been made.

### Creating a Radio