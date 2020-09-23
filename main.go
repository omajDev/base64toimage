package main

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	data, err := ioutil.ReadFile("token.txt")
	if err != nil {
		log.Fatal(err)
	}

	reader := base64.NewDecoder(base64.StdEncoding, bytes.NewReader(data))
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	outfilename := "drunk_gopher.png"
	f, err := os.OpenFile(outfilename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(f, m)
	if err != nil {
		log.Fatal(err)
	}

}
