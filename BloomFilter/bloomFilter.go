// Package bloomfilter
package bloomfilter

import (
	"github.com/spaolacci/murmur3"
)

type BloomFilter struct {
	rounds   uint32
	bitArray []bool
}

func NewBloomFilter(size int, rounds uint32) BloomFilter {
	var bloomFilter BloomFilter
	bloomFilter.bitArray = make([]bool, size)
	bloomFilter.rounds = rounds
	return bloomFilter
}

func (b *BloomFilter) AddKey(key string) {
	for i := range b.rounds {
		hash := murmur3.Sum32WithSeed([]byte(key), i)
		idx := hash % uint32(len(b.bitArray))
		b.bitArray[idx] = true
	}
}

func (b *BloomFilter) LookupKey(key string) bool {
	for i := range b.rounds {
		hash := murmur3.Sum32WithSeed([]byte(key), i)
		idx := hash % uint32(len(b.bitArray))
		if !b.bitArray[idx] {
			return false
		}
	}
	return true
}
