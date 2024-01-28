# Simple bencoder written in Go.

Usage: 
```golang
torrFile, err := os.Open("file.torrent")
if err != nil {
  panic("Cannot find specified file!")
}

reader := bufio.NewReader(torrFile)
output, err := decoder.Decode(reader)
```

To unmarshal the output use this: https://pkg.go.dev/github.com/mitchellh/mapstructure
