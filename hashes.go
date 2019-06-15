package go_checksum_benchmarks

import (
	"hash"
	"hash/crc64"
	"hash/fnv"

	"github.com/dgryski/go-spooky"
	"github.com/spaolacci/murmur3"
)

func NewCRC_64ISO() hash.Hash64 {
	return crc64.New(crc64.MakeTable(crc64.ISO))
}
func NewCRC_64ECMA() hash.Hash64 {
	return crc64.New(crc64.MakeTable(crc64.ECMA))
}
func NewFNV1_64() hash.Hash64 {
	return fnv.New64()
}
func NewFNV1_64a() hash.Hash64 {
	return fnv.New64a()
}

func NewMurmur3_64() hash.Hash64 {
	return murmur3.New64()
}

func NewJenkins_64() hash.Hash64 {
	return &spooky.Spooky{}
}
