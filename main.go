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
	sr := io.NewSectionReader(r, 0, 1<<63-1)

	var magic [4]byte
	if _, err := r.ReadAt(magic[0:], 0); err != nil {
		return err
	}

	combined := uint32(magic[0])<<24 | uint32(magic[1])<<16 |
		uint32(magic[2])<<8 | uint32(magic[3])

	var order binary.ByteOrder
	is64 := false
	switch combined {
	case macho.Magic32:
		fmt.Println("32-bit, big-endian")
		order = binary.BigEndian
	case macho.Magic64:
		fmt.Println("64-bit, big-endian")
		is64 = true
		order = binary.BigEndian
	case macho.Cigam32:
		fmt.Println("32-bit, little-endian")
		order = binary.LittleEndian
	case macho.Cigam64:
		fmt.Println("64-bit, little-endian")
		is64 = true
		order = binary.LittleEndian
	default:
		return fmt.Errorf("invalid magic number 0x%x", magic)
	}

	header := new(macho.FileHeader)
	if err := binary.Read(sr, order, header); err != nil {
		return err
	}

	switch header.CpuType {
	case macho.Cpu386:
		fmt.Println("CPU Type: x86")
	case macho.CpuAmd64:
		fmt.Println("CPU Type: x64")
	}

	switch header.Type {
	case macho.TypeObject:
		fmt.Println("relocatable object file")
	case macho.TypeExecutable:
		fmt.Println("executable file")
	case macho.TypeDylib:
		fmt.Println("dynamically bound shared library")
	case macho.TypeBundle:
		fmt.Println("dynamically bound bundle file")
	}

	fmt.Printf("%d load commands, totalling %d bytes\n",
		header.Ncmds, header.SizeOfCmds)

	offset := int64(macho.FileHeaderSize32)
	if is64 {
		offset = macho.FileHeaderSize64
	}

	fmt.Println(offset)
	return nil
}
