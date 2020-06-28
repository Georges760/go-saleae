package saleae

import (
	"encoding/csv"
	"io"
	"strconv"
	"strings"
)

type SPIExchange struct {
	Second float64
	Mosi   uint8
	Miso   uint8
}

func extractValue(s, sub string) (ret uint8, err error) {
	parts := strings.Split(s, ";")
	for _, p := range parts {
		if strings.Contains(p, sub) {
			sps := strings.Split(strings.TrimSpace(p), " ")
			ui, err := strconv.ParseUint(sps[1], 0, 8)
			if err == nil {
				ret = uint8(ui)
			}
			return ret, err
		}
	}
	return
}

func ParseCSV(file io.Reader) (spi []SPIExchange, err error) {
	// Read CSV File
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err == nil {
		// Parse it
		for _, r := range records {
			if r[1] == "SPI" {
				s := SPIExchange{}
				s.Second, _ = strconv.ParseFloat(r[0], 64)
				s.Mosi, _ = extractValue(r[2], "MOSI:")
				s.Miso, _ = extractValue(r[2], "MISO:")
				spi = append(spi, s)
			}
		}
	}
	return
}
