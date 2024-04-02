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

type SuffixStrings struct {
	Short string
	Long  string
}

var suffixStrings = map[Suffix]SuffixStrings{
	B:  {"B", "bytes"},
	KB: {"KB", "kilobytes"},
	MB: {"MB", "megabytes"},
	GB: {"GB", "gigabytes"},
	TB: {"TB", "terabytes"},
	PB: {"PB", "petabytes"},
}

func GetSuffixString(sfx Suffix) SuffixStrings { return suffixStrings[sfx] }

func GetSuffixByString(s string) *Suffix {
	for k, v := range suffixStrings {
		if s == v.Short {
			return &k
		}
	}
	return nil
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
	getSuffix := suffixStrings[Suffix(int(math.Floor(base)))]
	return strconv.FormatFloat(getSize, 'f', -1, 64) + " " + getSuffix.Short
}

func ConvertBytes(iIn int64, s Suffix) float64 {
	numLoops := getNumberOfLoops(s)
	f := float64(iIn)
	var multiplier float64 = 1024.00
	for i := 0; i < numLoops; i++ {
		f = f / multiplier
	}
	//base := math.Log(float64(iIn)) / math.Log(1024)
	//return round(math.Pow(1024, base-math.Floor(base)), .5, 2)
	return f
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

func GetValueTable(bytes int64) map[Suffix]float64 {
	valueTable := map[Suffix]float64{}
	valueTable[B] = float64(bytes)
	valueTable[KB] = ConvertBytes(bytes, KB)
	valueTable[MB] = ConvertBytes(bytes, MB)
	valueTable[GB] = ConvertBytes(bytes, GB)
	valueTable[TB] = ConvertBytes(bytes, TB)
	valueTable[PB] = ConvertBytes(bytes, PB)
	return valueTable
}
