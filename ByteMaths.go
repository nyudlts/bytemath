package ByteMaths

/* based on gist: https://gist.githubusercontent.com/anikitenko/ */

import (
	"math"
	"strconv"
)

var (
	sizeInMB float64 = 999 // This is in megabytes
	suffixes [5]string
)

type Suffix int

const (
	B Suffix = iota
	KB
	MB
	GB
	TB
)

func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}
func ConvertToBytes(count float64, suffix Suffix) float64 {

	c := count
	for i := 0; i < GetNumberOfLoops(suffix); i++ {
		c = c * 1024.0
	}

	return c
}

func ToHuman(size float64) string {
	suffixes[0] = "B"
	suffixes[1] = "KB"
	suffixes[2] = "MB"
	suffixes[3] = "GB"
	suffixes[4] = "TB"

	base := math.Log(size) / math.Log(1024)
	getSize := Round(math.Pow(1024, base-math.Floor(base)), .5, 2)
	getSuffix := suffixes[int(math.Floor(base))]
	return strconv.FormatFloat(getSize, 'f', -1, 64) + " " + string(getSuffix)
}

func GetNumberOfLoops(s Suffix) int {
	i := 0
	switch s {
	case B:
		i = 0
	case KB:
		i = 1
	case MB:
		i = 2
	case GB:
		i = 3
	case TB:
		i = 4
	}

	return i
}
