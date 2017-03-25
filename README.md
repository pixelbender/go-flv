# Golang: FLV File Format

[![Build Status](https://travis-ci.org/pixelbender/go-flv.svg)](https://travis-ci.org/pixelbender/go-flv)
[![Coverage Status](https://coveralls.io/repos/github/pixelbender/go-flv/badge.svg?branch=master)](https://coveralls.io/github/pixelbender/go-flv?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/pixelbender/go-flv)](https://goreportcard.com/report/github.com/pixelbender/go-flv)
[![GoDoc](https://godoc.org/github.com/pixelbender/go-flv?status.svg)](https://godoc.org/github.com/pixelbender/go-flv)

## Features

- [x] Header/Tag
- [ ] AAC/AVC Configuration
- [ ] MetaData
- [ ] AVC SPS/PPS Data
- [ ] AAC/AVC Bitstream

## Installation

```sh
go get github.com/pixelbender/go-flv
```

## FLV Reader

```go
package main

import (
    "github.com/pixelbender/go-flv/flv"
    "os"
    "log"
)

func main() {
	f, err := os.Open("example.flv")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    r := flv.NewReader(f)
    if err != nil {
        log.Fatal(err)
    }
    h, err := r.ReadHeader()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(h)
	for {
		tag, _, err := r.ReadTag()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		log.Println(tag)
	}
}
```

## Specifications

- [FLV: Video File Format Specification v10](https://www.adobe.com/content/dam/Adobe/en/devnet/flv/pdfs/video_file_format_spec_v10.pdf)
- [FLV: Adobe Flash Video File Format Specification v10.1](http://download.macromedia.com/f4v/video_file_format_spec_v10_1.pdf)
