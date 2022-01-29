OsrParser
----

[![Version 1.0.0](https://img.shields.io/badge/stable-1.0.0-brightgreen.svg "Version 1.0.0")](https://github.com/fauzanardh/OsrParser) [![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT) [![Supported go version: 1.11](https://img.shields.io/badge/go-1.11-green.svg "Supported go versions: 1.11")](https://golang.org/dl/)


Osu! Replay parser build with GO version 1.11 (modified for personal use)


Refrence
----

- BinaryReader : [Mempler's Repo](https://github.com/Mempler/osubinary/)
- ULEB128      : [Bnch's Repo](https://github.com/bnch/uleb128/)

Usage
----

 **First of all**
 install lzma package from [Itchio's Repo](https://github.com/itchio/lzma)

 `go get github.com/itchio/lzma`

 and then install this package

 `go get github.com/fauzanardh/OsrParser`

 (optional) add your osu!api apikey to get the beatmap data, edit the APIKEY from `$GOPATH/src/github.com/fauzanardh/OsrParser/constants.go`

 **How to use**
 ```go
 package main

 import (
 	"io/ioutil"

 	"github.com/fauzanardh/OsrParser"
 )

 func main() {
 	buffer, _ := ioutil.ReadFile("rep.osr")
 	replay, _ := OsrParser.ParseReplay(buffer, false)
 	replay.PrintReplay()
 	replay.PrintAllBeatmap()
 }
 ```

 Screenshot
 ----
 ![result_exmaple](example.png)

Author
----
[@Fauzanardh](https://github.com/fauzanardh)
