// Copyright 2016 ego authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

/* Package murmur Murmur3 hash function based on
http://en.wikipedia.org/wiki/MurmurHash
*/

package murmur

import (
	"reflect"
	"unsafe"
)

const (
	c1 = 0xcc9e2d51
	c2 = 0x1b873593
	c3 = 0x85ebca6b
	c4 = 0xc2b2ae35
	nh = 0xe6546b64
)

var (
	// defaultSeed default murmur seed
	defaultSeed = uint32(1)
)

// Murmur3 returns a hash from the provided key using the specified seed.
func Murmur3(key []byte, seed ...uint32) uint32 {
	return Sum32(*(*string)((unsafe.Pointer)(
		&reflect.StringHeader{
			Len:  len(key),
			Data: (*reflect.SliceHeader)(unsafe.Pointer(&key)).Data,
		})),
		seed...)

	// return Sum32(string(key), seed...)
}

// Sum32 returns a hash from the provided key using the seed.
func Sum32(key string, seed ...uint32) (hash uint32) {
	hash = defaultSeed
	if len(seed) > 0 {
		hash = seed[0]
	}

	var nbytes = len(key) / 4 * 4
	for i := 0; i < nbytes; i += 4 {
		k := uint32(key[i]) | uint32(key[i+1])<<8 |
			uint32(key[i+2])<<16 | uint32(key[i+3])<<24

		k *= c1
		k = (k << 15) | (k >> 17)
		k *= c2
		hash ^= k

		hash = (hash << 13) | (hash >> 19)
		hash = hash*5 + nh
	}

	var remaining uint32
	switch len(key) & 3 {
	case 3:
		remaining += uint32(key[nbytes+2]) << 16
		fallthrough
	case 2:
		remaining += uint32(key[nbytes+1]) << 8
		fallthrough
	case 1:
		remaining += uint32(key[nbytes])
		remaining *= c1

		remaining = (remaining << 15) | (remaining >> 17)
		remaining = remaining * c2
		hash ^= remaining
	}

	hash ^= uint32(len(key))
	hash ^= hash >> 16
	hash *= c3
	hash ^= hash >> 13
	hash *= c4
	hash ^= hash >> 16

	return
}

// func Sum64() {
// }

// func Sum128() {
// }
