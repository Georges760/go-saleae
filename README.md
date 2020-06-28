# go-saleae : Go Library to interface with Saleae Protocol Analyzer

## Overview [![GoDoc](https://godoc.org/github.com/Georges760/go-saleae?status.svg)](https://godoc.org/github.com/Georges760/go-saleae)

## Install

```
go get github.com/Georges760/go-saleae
```

## Example

```
package main

import (
	"flag"
	"fmt"
	"github.com/Georges760/go-saleae"
	"log"
	"os"
)

func main() {
	// Check args
	pfile := flag.String("f", "", "CSV file from Saleae Decoded Protocol export")
	flag.Parse()
	if *pfile == "" {
		log.Fatal("-f arg is mandatory")
	}
	// Open CSV file
	recordFile, err := os.Open(*pfile)
	if err != nil {
		log.Fatal("Error opening CSV file:", err)
	}
	// Parse it
	spi, err := saleae.ParseCSV(recordFile)
	if err != nil {
		log.Fatal("Error parsing CSV file:", err)
	}
	// Print result
	fmt.Println(spi)
	// Exit
	os.Exit(0)
}
```

## License

MIT.
