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
	bytesPerSectorOffset     = 0x0B // 2字节
	sectorsPerClusterOffset  = 0x0D // 1字节
	reservedSectorsOffset    = 0x0E // 2字节，相对于第0个扇区来说，不是FAT32的第0个扇区
	fatTableCountOffset      = 0x10 // 1字节
	hidenSectorCountOffset   = 0x1C // 4字节
	fsTotalSectorCountOffset = 0x20 // 4字节
	sectorsPerFatTableOffset = 0x24 // 4字节，每个FAT表占用的扇区数
	rootClusterNumOffset     = 0x2C // 4字节，根目录所在第一个簇的簇号
	fsInfoSectorNumOffset    = 0x30 // 2字节，FSINFO扇区号
	labelOffset              = 0x47 // 11字节

	labelLength = 11
)

type DBR struct {
	BytesPerSector     uint16
	SectorsPerCluster  uint8
	ReservedSectors    uint16
	FatTableCount      uint8
	HidenSectorCount   uint32
	FsTotalSectorCount uint32
	SectorsPerFatTable uint32
	RootClusterNum     uint32
	FsInfoSectorNum    uint16
	Label              string
}

func (d *DBR) Read(sector []byte) {
	d.BytesPerSector = bin.Uint16(sector[bytesPerSectorOffset:])
	d.SectorsPerCluster = uint8(sector[sectorsPerClusterOffset])
	d.ReservedSectors = bin.Uint16(sector[reservedSectorsOffset:])
	d.FatTableCount = uint8(sector[fatTableCountOffset])
	d.HidenSectorCount = bin.Uint32(sector[hidenSectorCountOffset:])
	d.FsInfoSectorNum = uint16(bin.Uint32(sector[fsInfoSectorNumOffset:]))
	d.SectorsPerFatTable = bin.Uint32(sector[sectorsPerFatTableOffset:])
	d.RootClusterNum = bin.Uint32(sector[rootClusterNumOffset:])
	d.FsInfoSectorNum = bin.Uint16(sector[fsInfoSectorNumOffset:])
	d.Label = string(sector[labelOffset : labelOffset+labelLength])
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
