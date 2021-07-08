/*
 * ------------------------------------------------------------------------
 *   File Name: dbr.go
 *      Author: Zhao Yanbai
 *              2021-07-07 19:18:25 Wednesday CST
 * Description: none
 * ------------------------------------------------------------------------
 */

package fat32

import (
	"fmt"
	"strings"
)

const (
	offsetBytesPerSector     = 0x0B // 2字节
	offsetSectorsPerCluster  = 0x0D // 1字节
	offsetReservedSectors    = 0x0E // 2字节，相对于第0个扇区来说，不是FAT32的第0个扇区
	offsetFatTableCount      = 0x10 // 1字节
	offsetHidenSectorCount   = 0x1C // 4字节
	offsetFsTotalSectorCount = 0x20 // 4字节
	offsetSectorsPerFatTable = 0x24 // 4字节，每个FAT表占用的扇区数
	offsetRootClusterNum     = 0x2C // 4字节，根目录所在第一个簇的簇号
	offsetFsInfoSectorNum    = 0x30 // 2字节，FSINFO扇区号
	offsetLabel              = 0x47 // 11字节

	labelLength = 11
)

type DBR struct {
	BytesPerSector     Int
	SectorsPerCluster  Int
	ReservedSectors    Int
	FatTableCount      Int
	HidenSectorCount   Int
	FsTotalSectorCount Int
	SectorsPerFatTable Int
	RootClusterNum     Int
	FsInfoSectorNum    Int
	Label              string
}

func (d *DBR) Read(sector []byte) {
	d.BytesPerSector = Bin.Uint16(sector[offsetBytesPerSector:])
	d.SectorsPerCluster = Int(sector[offsetSectorsPerCluster])
	d.ReservedSectors = Bin.Uint16(sector[offsetReservedSectors:])
	d.FatTableCount = Int(sector[offsetFatTableCount])
	d.HidenSectorCount = Bin.Uint32(sector[offsetHidenSectorCount:])
	d.FsTotalSectorCount = Bin.Uint32(sector[offsetFsTotalSectorCount:])
	d.SectorsPerFatTable = Bin.Uint32(sector[offsetSectorsPerFatTable:])
	d.RootClusterNum = Bin.Uint32(sector[offsetRootClusterNum:])
	d.FsInfoSectorNum = Bin.Uint16(sector[offsetFsInfoSectorNum:])
	d.Label = string(sector[offsetLabel : offsetLabel+labelLength])
	d.Label = strings.TrimSuffix(d.Label, " ")
}

func (d DBR) String() string {
	s := ""
	s += fmt.Sprintf("每扇区字节数: %v\n", d.BytesPerSector)
	s += fmt.Sprintf("每簇扇区数: %v\n", d.SectorsPerCluster)
	s += fmt.Sprintf("保留扇区数: 0x%08X %v\n", d.ReservedSectors, d.ReservedSectors)
	s += fmt.Sprintf("Fat表个数: %v\n", d.FatTableCount)
	s += fmt.Sprintf("隐藏扇区数: 0x%08x %v\n", d.HidenSectorCount, d.HidenSectorCount)
	s += fmt.Sprintf("文件系统总扇区数: %d\n", d.FsTotalSectorCount)
	s += fmt.Sprintf("每个Fat表的扇区数: %v\n", d.SectorsPerFatTable)
	s += fmt.Sprintf("Root目录所在的第一个簇号: %v\n", d.RootClusterNum)
	s += fmt.Sprintf("FSINFO的扇区号: %v\n", d.FsInfoSectorNum)
	s += fmt.Sprintf("卷标: %v\n", d.Label)
	return s
}
