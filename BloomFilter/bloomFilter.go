// Package bloomfilter
package bloomfilter

import (
	"math"

	"github.com/spaolacci/murmur3"
)

type BloomFilter struct {
	rounds   uint32
	bitArray []byte
}

func NewBloomFilter(size int, rounds uint32) BloomFilter {
	var bloomFilter BloomFilter
	bloomFilter.bitArray = make([]byte, int(math.Ceil(float64(size)/8)))
	bloomFilter.rounds = rounds
	return bloomFilter
}

func (b *BloomFilter) AddKey(key string) {
	for i := range b.rounds {
		hash := murmur3.Sum32WithSeed([]byte(key), i)
		idx := hash % uint32(len(b.bitArray)*8)
		byteIdx := int(math.Floor(float64(idx) / 8))
		bitPos := idx % 8
		tmp := byte(1 << (8 - bitPos - 1))
		b.bitArray[byteIdx] |= tmp
	}
}

func (b *BloomFilter) LookupKey(key string) bool {
	for i := range b.rounds {
		hash := murmur3.Sum32WithSeed([]byte(key), i)
		idx := hash % uint32(len(b.bitArray)*8)
		byteIdx := int(math.Floor(float64(idx) / 8))
		bitPos := idx % 8
		tmp := byte(1 << (8 - bitPos - 1))
		if (b.bitArray[byteIdx] & tmp) == byte(0) {
			return false
		}
	}
	return true
}
