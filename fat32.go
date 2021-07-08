/*
 * ------------------------------------------------------------------------
 *   File Name: fat32.go
 *      Author: Zhao Yanbai
 *              2021-07-07 23:28:03 Wednesday CST
 * Description: none
 * ------------------------------------------------------------------------
 */

package fat32

import "encoding/binary"

type Fat32 struct {
	Mbr MBR
	Dbr DBR
}

type Offset uint64

func (f Fat32) GetFATAreaOffset() Offset {
	sectorSize := f.Dbr.BytesPerSector

	offset := Offset(f.Mbr.PTE.StartLBA)
	offset += Offset(f.Dbr.ReservedSectors)

	return offset * Offset(sectorSize)
}

func (f Fat32) GetDataAreaOffset() Offset {
	offset := f.GetFATAreaOffset()
	offset += Offset(f.Dbr.SectorsPerFatTable) * Offset(f.Dbr.FatTableCount) * Offset(f.Dbr.BytesPerSector)
	return offset
}

var bin = binary.LittleEndian
