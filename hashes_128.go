package go_checksum_benchmarks

import (
	"encoding/hex"
	"hash"
	"hash/fnv"

	"github.com/dgryski/go-spooky"
	"github.com/spaolacci/murmur3"
)

// Hash128 is the common interface implemented by all 128-bit hash functions.
type Hash128 interface {
	hash.Hash
	Sum128Bytes() []byte
}

type fnv1wrapper struct {
	hash.Hash
}

func (w *fnv1wrapper) Sum128Bytes() []byte {
	return w.Sum(nil)
}

func NewFNV1_128() Hash128 {
	return &fnv1wrapper{fnv.New128()}
}
func NewFNV1_128a() Hash128 {
	return &fnv1wrapper{fnv.New128a()}
}

type murmur3wrapper struct {
	murmur3.Hash128
}

func (w *murmur3wrapper) Sum128Bytes() []byte {
	a, b := w.Sum128()
	return uintToBytes([2]uint64{a, b})
}

func NewMurmur3_128() Hash128 {
	return &murmur3wrapper{murmur3.New128()}
}

type jenkinswrapper struct {
	hash.Hash
}

func (w *jenkinswrapper) Sum128Bytes() []byte {
	return w.Sum(nil)
}
func NewJenkins_128() Hash128 {
	return &jenkinswrapper{&spooky.Spooky{}}
}

func uintToBytes(s [2]uint64) []byte {
	return []byte{
		byte(s[0] >> 56), byte(s[0] >> 48), byte(s[0] >> 40), byte(s[0] >> 32), byte(s[0] >> 24), byte(s[0] >> 16), byte(s[0] >> 8), byte(s[0]),
		byte(s[1] >> 56), byte(s[1] >> 48), byte(s[1] >> 40), byte(s[1] >> 32), byte(s[1] >> 24), byte(s[1] >> 16), byte(s[1] >> 8), byte(s[1]),
	}
}

func Hex(h Hash128) string {
	return hex.EncodeToString(h.Sum128Bytes())
}
