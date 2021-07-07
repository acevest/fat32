/*
 * ------------------------------------------------------------------------
 *   File Name: info.go
 *      Author: Zhao Yanbai
 *              2021-07-07 17:26:21 Wednesday CST
 * Description: none
 * ------------------------------------------------------------------------
 */

package main

import (
	"bufio"
	"encoding/binary"
	"fat32"
	"flag"
	"log"
	"os"
)

func main() {
	defer log.Println("Program Exited...")

	var path string
	flag.StringVar(&path, "p", "", "file path")
	flag.Parse()

	if path == "" {
		log.Fatalf("please specify file path")
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("open file: %v error %v", path, err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	sector := make([]byte, 512)
	n, err := reader.Read(sector)
	if err != nil {
		log.Fatalf("read MBR fail: %v", err)
	}
	if n != len(sector) {
		log.Fatalf("not 512 bytes")
	}

	var mbr fat32.MBR
	mbr.Read(sector)

	log.Printf("%x", mbr.PTEs[0].StartLBA*512)

	pos, err := file.Seek(int64(mbr.PTEs[0].StartLBA*512), 0)
	if err != nil {
		log.Printf("read DBR error: %v", err)
	}
	log.Printf("pos : %v", pos)
	reader = bufio.NewReader(file)
	n, err = reader.Read(sector)
	if err != nil {
		log.Fatalf("read DBR fail: %v", err)
	}
	if n != len(sector) {
		log.Fatalf("not 512 bytes")
	}

	var dbr fat32.DBR

	dbr.Read(sector)

	log.Printf("%x", sector)

	log.Printf("DBR: \n%v", dbr)
	// for i := 0; i < fat32.GPTEntryCount; i++ {
	// 	gptType := fat32.PartitionType(bin.Uint16(sector[fat32.GPTPos+i*fat32.PTESize+fat32.PTETypeOffset:]))
	// 	lba := bin.Uint32(sector[fat32.GPTPos+i*fat32.PTESize+fat32.PTEStartLBAOffset:])
	// 	totalSectors := bin.Uint32(sector[fat32.GPTPos+i*fat32.PTESize+fat32.PTETotalSectorsOffset:])
	// 	log.Printf("PartitionTableEntry[%d] type: %v LBA: 0x%08X %d total sectors: 0x%08X %d", i, gptType, lba, lba, totalSectors, totalSectors)
	// }

	//log.Printf("%x", sector)
}

var bin = binary.LittleEndian
