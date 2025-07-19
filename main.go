package main

import (
	bloomfilter "Kaushik1766/BloomFilter/BloomFilter"
	"fmt"
)

func main() {
	b := bloomfilter.NewBloomFilter(20, 3)
	keys := []string{"hello", "hi", "asdfsda", "kkde"}
	for _, key := range keys {
		b.AddKey(key)
	}

	fmt.Println(b.LookupKey("hello world"))
	// for _, key := range keys {
	// 	fmt.Println(b.LookupKey(key))
	// }
}
