package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/croese/rambo/macho"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("USAGE: rambo <path-to-mach-o-binary>\n")
		os.Exit(1)
	}

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = dumpDetails(file)
	if err != nil {
		log.Fatal(err)
	}
}

func dumpDetails(r io.ReaderAt) error {
	var magic [4]byte
	if _, err := r.ReadAt(magic[0:], 0); err != nil {
		return err
	}

	be := binary.BigEndian.Uint32(magic[0:])
	le := binary.LittleEndian.Uint32(magic[0:])

	switch {
	case be == macho.Magic32:
		fmt.Println("32-bit, big-endian")
	case be == macho.Magic64:
		fmt.Println("64-bit, big-endian")
	case le == macho.Cigam32:
		fmt.Println("32-bit, little-endian")
	case be == macho.Cigam64:
		fmt.Println("64-bit, little-endian")
	}
	return nil
}
