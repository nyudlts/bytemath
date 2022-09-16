package bytemath

/* based on gist: https://gist.github.com/anikitenko/b41206a49727b83a530142c76b1cb82d */

import (
	"math"
	"strconv"
)

type Suffix int

const (
	B Suffix = iota
	KB
	MB
	GB
	TB
	PB
)

var suffixes = map[Suffix]string{
	B:  "B",
	KB: "KB",
	MB: "MB",
	GB: "GB",
	TB: "TB",
	PB: "PB",
}

func ConvertToBytes(bytes float64, suffix Suffix) float64 {
	for i := 0; i < getNumberOfLoops(suffix); i++ {
		bytes = bytes * 1024.0
	}
	return bytes
}

func ConvertBytesToHumanReadable(sizeInBytes int64) string {
	floatSize := float64(sizeInBytes)
	base := math.Log(floatSize) / math.Log(1024)
	getSize := round(math.Pow(1024, base-math.Floor(base)), .5, 2)
	getSuffix := suffixes[Suffix(int(math.Floor(base)))]
	return strconv.FormatFloat(getSize, 'f', -1, 64) + " " + string(getSuffix)
}

func ConvertBytes(iIn int64, s Suffix) float64 {
	numLoops := getNumberOfLoops(s)
	fIn := float64(iIn)
	var multiplier float64 = 1024
	for i := 0; i < numLoops; i++ {
		fIn = fIn * multiplier
	}
	base := math.Log(float64(iIn)) / math.Log(1024)
	return round(math.Pow(1024, base-math.Floor(base)), .5, 2)
}

func round(val float64, roundOn float64, places int) (newVal float64) {
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

func getNumberOfLoops(s Suffix) int {
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
	case PB:
		i = 5
	}

	return i
}
