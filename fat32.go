/*
 * ------------------------------------------------------------------------
 *   File Name: fat32.go
 *      Author: Zhao Yanbai
 *              2021-07-07 23:28:03 Wednesday CST
 * Description: none
 * ------------------------------------------------------------------------
 */

package fat32

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Fat32 struct {
	File *os.File
	mbr  MBR
	dbr  DBR

	fatTable []byte
}

type Offset uint64
type Int uint64

func (f Fat32) GetFatTableAreaOffset() Offset {
	sectorSize := f.SectorSize()

	offset := Offset(f.mbr.PTE.StartLBA)
	offset += Offset(f.dbr.ReservedSectors)

	return offset * Offset(sectorSize)
}

func (f Fat32) GetDataAreaOffset() Offset {
	offset := f.GetFatTableAreaOffset()
	offset += Offset(f.dbr.SectorsPerFatTable) * Offset(f.dbr.FatTableCount) * Offset(f.SectorSize())
	return offset
}

func (f *Fat32) ReadFAT() error {
	var err error
	err = f.ReadMBR()
	if err != nil {
		return err
	}

	err = f.ReadDBR()
	if err != nil {
		return err
	}

	log.Printf("DBR: \n%v", f.dbr)

	err = f.ReadFatTable()
	if err != nil {
		return err
	}

	root, err := f.ReadCluster(2)
	if err != nil {
		return err
	}
	log.Printf("%x", root[:128])

	return nil
}

func (f *Fat32) ReadMBR() error {
	sector, err := f.ReadSector(0)
	if err != nil {
		return err
	}

	f.mbr.Read(sector)

	return nil
}

func (f *Fat32) ReadDBR() error {
	sector, err := f.ReadSector(Int(f.mbr.PTE.StartLBA))
	if err != nil {
		return err
	}

	//log.Printf("%x", sector)
	f.dbr.Read(sector)

	return nil
}

func (f *Fat32) ReadFatTable() error {
	size := f.dbr.SectorsPerFatTable * Int(f.SectorSize())
	offset := f.GetFatTableAreaOffset()

	//f.fatTable = make([]byte, size)

	var err error
	f.fatTable, err = f.ReadData(offset, size)

	return err
}

func (f Fat32) ReadData(offset Offset, size Int) ([]byte, error) {
	data := make([]byte, uint64(f.SectorSize()))
	n, err := f.File.ReadAt(data, int64(offset))
	if n == 0 {
		if err == io.EOF {
			return nil, err
		}

		if err == nil {
			return nil, fmt.Errorf("maybe buffer set to zero?")
		}

		return nil, err
	}

	if n != int(f.SectorSize()) {
		return nil, fmt.Errorf("read data failed, read %d/%d bytes", n, f.SectorSize())
	}

	return data, err
}

func (f Fat32) ReadSector(sectNumber Int) ([]byte, error) {
	offset := Offset(sectNumber) * Offset(f.SectorSize())

	return f.ReadData(offset, Int(f.SectorSize()))
}

func (f Fat32) ReadCluster(clusterNumber Int) ([]byte, error) {

	if clusterNumber < 2 {
		return nil, fmt.Errorf("can not read cluster %d", clusterNumber)
	}

	bytesPerCluster := f.dbr.SectorsPerCluster * f.dbr.BytesPerSector

	// offset := Offset(f.mbr.PTE.StartLBA)
	// offset += Offset(f.dbr.ReservedSectors)
	// offset += Offset(f.dbr.SectorsPerFatTable * f.dbr.FatTableCount)
	// offset *= Offset(f.SectorSize())

	offset := f.GetDataAreaOffset()
	offset += Offset(clusterNumber-2) * Offset(bytesPerCluster)

	return f.ReadData(offset, bytesPerCluster)
}

func (f Fat32) SectorSize() int {
	if f.dbr.BytesPerSector == 0 {
		return 512
	}

	return int(f.dbr.BytesPerSector)
}
