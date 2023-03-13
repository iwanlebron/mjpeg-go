package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"

	"github.com/ivanlebron/mjpeg-go"
)

func main() {
	checkErr := func(err error) {
		if err != nil {
			panic(err)
		}
	}

	outName := "gopher.mp4"
	aw, err := mjpeg.New(outName, 144, 108, 2)
	checkErr(err)

	// Create a movie from images:
	matches, err := filepath.Glob("*.jpg")
	checkErr(err)

	if matches == nil || len(matches) <= 0 {
		fmt.Printf("No images was found, exits directly")
		return
	}

	sort.Strings(matches)

	fmt.Println("Found images:", matches)
	for _, name := range matches {
		data, err := ioutil.ReadFile(name)
		checkErr(err)
		checkErr(aw.AddFrame(data))
	}

	checkErr(aw.Close())
	fmt.Printf("%s was written successfully.\n", outName)
}
