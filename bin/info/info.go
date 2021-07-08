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

	// sector := make([]byte, 512)
	// n, err := file.ReadAt(sector, 0)
	// if err != nil {
	// 	log.Fatalf("read MBR fail: %v", err)
	// }
	// if n != len(sector) {
	// 	log.Fatalf("not 512 bytes")
	// }

	fs := fat32.Fat32{
		File: file,
	}

	err = fs.ReadFAT()
	if err != nil {
		log.Fatalf("aaa %v", err)
	}

	//log.Printf("DBR: \n%v")

	// var mbr fat32.MBR
	// mbr.Read(sector)

	// log.Printf("%x", mbr.PTE.StartLBA*512)

	// pos, err := file.Seek(int64(mbr.PTE.StartLBA*512), 0)
	// if err != nil {
	// 	log.Printf("read DBR error: %v", err)
	// }
	// log.Printf("pos : %v", pos)
	// n, err = file.ReadAt(sector, 0)
	// if err != nil {
	// 	log.Fatalf("read DBR fail: %v", err)
	// }
	// if n != len(sector) {
	// 	log.Fatalf("not 512 bytes")
	// }

	// var dbr fat32.DBR

	// dbr.Read(sector)

	// log.Printf("%x", sector)

	// log.Printf("DBR: \n%v", dbr)
}

var bin = binary.LittleEndian
