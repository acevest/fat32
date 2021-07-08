/*
 * ------------------------------------------------------------------------
 *   File Name: bin.go
 *      Author: Zhao Yanbai
 *              2021-07-08 12:11:01 Thursday CST
 * Description: none
 * ------------------------------------------------------------------------
 */

package fat32

import "encoding/binary"

var bin = binary.LittleEndian

var Bin LittleEndian

type LittleEndian struct{}

func (LittleEndian) Uint16(b []byte) Int {
	return Int(binary.LittleEndian.Uint16(b))
}

func (LittleEndian) Uint32(b []byte) Int {
	return Int(binary.LittleEndian.Uint32(b))
}

func (LittleEndian) Uint64(b []byte) Int {
	return Int(binary.LittleEndian.Uint64(b))
}

func (LittleEndian) String() string {
	return "FAT32.LittleEndian"
}
