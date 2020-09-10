murmur
======
[![CircleCI Status](https://circleci.com/gh/vcaesar/murmur.svg?style=shield)](https://circleci.com/gh/vcaesar/murmur)
[![codecov](https://codecov.io/gh/vcaesar/murmur/branch/master/graph/badge.svg)](https://codecov.io/gh/vcaesar/murmur)
[![Build Status](https://travis-ci.org/vcaesar/murmur.svg)](https://travis-ci.org/vcaesar/murmur)
[![Go Report Card](https://goreportcard.com/badge/github.com/vcaesar/murmur)](https://goreportcard.com/report/github.com/vcaesar/murmur)
[![GoDoc](https://godoc.org/github.com/vcaesar/murmur?status.svg)](https://godoc.org/github.com/vcaesar/murmur)
[![Release](https://github-release-version.herokuapp.com/github/vcaesar/murmur/release.svg?style=flat)](https://github.com/vcaesar/murmur/releases/latest)
[![Join the chat at https://gitter.im/vcaesar/ego](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/vcaesar/ego?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

Go Murmur3 hash implementation

## Installing
```Go
go get -u github.com/vcaesar/murmur
```

## Use

```Go
package main

import (
	"log"

	"github.com/vcaesar/murmur"
)

func main() {
	var str = "github.com"

	sum32 := murmur.Sum32(str)
	log.Println("hash32...", sum32)

	sum32 = murmur.Sum32(str, 0)
	log.Println("hash32...", hash32)

	hash32 := murmur.Murmur3([]byte(str))
	log.Println("hash32...", hash32)

	hash32 = murmur.Murmur3([]byte(str), 1)
	log.Println("hash32...", hash32)
}
```

Based on [MurmurHash](http://en.wikipedia.org/wiki/MurmurHash), thanks murmur3.