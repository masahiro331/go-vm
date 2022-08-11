package main

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"strings"

	"github.com/masahiro331/go-vm"
)

func main() {
	f, _ := os.Open(os.Args[1])

	filesystem, err := vm.Open(f, "Linux")
	if err != nil {
		log.Fatalf("%+v", err)
	}
	err = fs.WalkDir(filesystem, "etc", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatalf("%+v", err)
		}
		if d.IsDir() {
			return nil
		}
		if strings.Contains(path, "os-release") {
			fmt.Println(path)
			f, err := filesystem.Open(path)
			if err != nil {
				log.Fatal(err)
			}
			data, err := io.ReadAll(f)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(data))
		}
		return nil
	})

	if err != nil {
		log.Fatalf("%+v", err)
	}
}