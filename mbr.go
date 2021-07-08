/*
 * ------------------------------------------------------------------------
 *   File Name: mbr.go
 *      Author: Zhao Yanbai
 *              2021-07-07 18:00:12 Wednesday CST
 * Description: none
 * ------------------------------------------------------------------------
 */

package fat32

import (
	"fmt"
	"log"
)

const (
	SectorSize    = 512
	GPTPos        = 0x01BE
	GPTEntryCount = 4
)

type MBR struct {
	PTE  PartitionTableEntry
	ptes []PartitionTableEntry
}

func (m *MBR) Read(sector []byte) {
	for i := 0; i < GPTEntryCount; i++ {
		var pte PartitionTableEntry
		pte.Type = PartitionType(bin.Uint16(sector[GPTPos+i*PTESize+PTETypeOffset:]))
		pte.StartLBA = bin.Uint32(sector[GPTPos+i*PTESize+PTEStartLBAOffset:])
		pte.TotalSectors = bin.Uint32(sector[GPTPos+i*PTESize+PTETotalSectorsOffset:])

		ignore := ""
		if pte.Type == PartitionTypeFat32 {
			m.ptes = append(m.ptes, pte)
			ignore = "[KEEPED]"
			m.PTE = pte
		} else {
			ignore = "[ignored]"
		}

		log.Printf("PartitionTableEntry[%d] %v %v", i, pte, ignore)
	}
}

const (
	PTESize               = 16
	PTETypeOffset         = 4
	PTEStartLBAOffset     = 8
	PTETotalSectorsOffset = 12
)

type PartitionTableEntry struct {
	State        uint8
	StartHead    uint8
	StartSC      SectorCylinder
	Type         PartitionType
	EndHead      uint8
	EndSC        SectorCylinder
	StartLBA     uint32
	TotalSectors uint32
}

func (pte PartitionTableEntry) String() string {
	return fmt.Sprintf(" type: %10v LBA: 0x%08X %10d total sectors: 0x%08X %10d", pte.Type, pte.StartLBA, pte.StartLBA, pte.TotalSectors, pte.TotalSectors)
}

// GPT: GUID(Globals Unique Identifiers) Partition Table
type PartitionType uint8

type SectorCylinder uint16

const (
	PartitionTypeUnused   PartitionType = 0x00
	PartitionTypeFat16    PartitionType = 0x06
	PartitionTypeFat32    PartitionType = 0x0B
	PartitionTypeFat32LBA PartitionType = 0x0C
	PartitionTypeExtend   PartitionType = 0x05
	PartitionTypeNTFS     PartitionType = 0x07
	PartitionTypeLBAMode  PartitionType = 0x0F
	PartitionTypeLinux    PartitionType = 0x83
)

var PartitionTypeMap = map[PartitionType]string{
	PartitionTypeUnused:   "Unused",
	PartitionTypeFat16:    "Fat16",
	PartitionTypeFat32:    "Fat32",
	PartitionTypeFat32LBA: "Fat32(LBA)",
	PartitionTypeExtend:   "Extend",
	PartitionTypeNTFS:     "NTFS",
	PartitionTypeLBAMode:  "LBAMode",
	PartitionTypeLinux:    "Linux",
}

func (g PartitionType) String() string {
	if s, ok := PartitionTypeMap[g]; ok {
		return s
	}

	return "NULL"
}
