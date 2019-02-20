package main

import (
	"fmt"
	"libs/KeyGenerator"
	"os"
	"strconv"
)

func printHelp() {
	fmt.Print("TODO, help")
}

type Args []string

func (a Args) find(s string) int {
	for i, v := range a {
		if v == s {
			return i
		}
	}
	return -1
}

func generateKeys(fn string, size int) {
	privateKey := KeyGenerator.GeneratePrivate(size)

	KeyGenerator.GeneratePEMFiles(privateKey, fn)
}

func main() {
	args := Args(os.Args[1:])

	size := 2048
	name := "out"

	if len(args) == 0 {
		printHelp()
	} else {
		if args.find("-gen") >= 0 {
			outFlagIndex := args.find("-o")
			sizeFlagIndex := args.find("-s")

			if sizeFlagIndex >= 0 && (sizeFlagIndex+1) <= len(args) {
				s, err := strconv.ParseInt(args[sizeFlagIndex+1], 10, 32)
				if err == nil {
					size = int(s)
				}
			}

			if outFlagIndex >= 0 && (outFlagIndex+1) <= len(args) {
				name = args[outFlagIndex+1]
			}

			fmt.Println("RSA Key pair generaiton")
			fmt.Println("Key size:", size)

			fmt.Println("In progress...")
			generateKeys(name, size)
			fmt.Println("Done")
			fmt.Println("Private out:", name+".key")
			fmt.Println("Public out:", name+".crt")

			os.Exit(0)
		}
	}
}
