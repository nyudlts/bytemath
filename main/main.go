package main

import (
	"flag"
	"fmt"
	"github.com/nyudlts/bytemath"
	"os"
	"strings"
)

var (
	suffix string
	sfx    *bytemath.Suffix
	dec    float64
)

func init() {
	flag.StringVar(&suffix, "suffix", "", "")
	flag.Float64Var(&dec, "decimal", float64(0), "")
}

func main() {
	flag.Parse()

	if dec == 0 {
		panic(fmt.Errorf("Cannot do any calculations with a value of 0"))
		os.Exit(1)
	}

	sfx = bytemath.GetSuffixByString(strings.ToUpper(suffix))
	if sfx == nil {
		panic(fmt.Errorf("Invalid Suffix, supported suffixes: B, KB, MB, GB, TB, and PB"))
		os.Exit(1)
	}

	printTable()
}

func printTable() {
	fmt.Printf("%f %s is:\n", dec, bytemath.GetSuffixString(*sfx).Short)
	b := bytemath.ConvertToBytes(dec, *sfx)
	bInt64 := int64(b)
	valueTable := bytemath.GetValueTable(bInt64)
	fmt.Printf("  %d %s\n", int64(valueTable[bytemath.B]), bytemath.GetSuffixString(bytemath.B).Long)
	fmt.Printf("  %.2f %s\n", valueTable[bytemath.KB], bytemath.GetSuffixString(bytemath.KB).Long)
	fmt.Printf("  %.2f %s\n", valueTable[bytemath.MB], bytemath.GetSuffixString(bytemath.MB).Long)
	fmt.Printf("  %.2f %s\n", valueTable[bytemath.GB], bytemath.GetSuffixString(bytemath.GB).Long)
	fmt.Printf("  %.2f %s\n", valueTable[bytemath.TB], bytemath.GetSuffixString(bytemath.TB).Long)
	fmt.Printf("  %.2f %s\n", valueTable[bytemath.PB], bytemath.GetSuffixString(bytemath.PB).Long)
}
