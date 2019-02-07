package main

import (
	"compress/gzip"
	"fmt"
	"libs/FileSystemSerializer"
)

type ss string

func (s ss) Write(b []byte) (int, error) {
	fmt.Print(b)
	return len(b), nil
}

func main() {
	var s ss

	compressor := gzip.NewWriter(s)
	defer compressor.Close()

	FileSystemSerializer.Serialize("../src", compressor)

}
