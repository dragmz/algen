package main

import (
	"encoding/csv"
	"flag"
	"os"
	"strings"

	"github.com/algorand/go-algorand-sdk/mnemonic"
	"github.com/dragmz/algen"
)

func main() {
	var startsWith string
	var endsWith string

	flag.StringVar(&startsWith, "starts-with", "", "generate address that starts with a given string")
	flag.StringVar(&endsWith, "ends-with", "", "generate address that ends with a given string")
	flag.Parse()

	a, err := algen.GenerateAddress(algen.GenerateArgs{
		StartsWith: strings.ToUpper(startsWith),
		EndsWith:   strings.ToUpper(endsWith),
	})
	if err != nil {
		panic(err)
	}

	m, err := mnemonic.FromPrivateKey(a.PrivateKey)
	if err != nil {
		panic(err)
	}

	w := csv.NewWriter(os.Stdout)
	w.Write([]string{a.Address.String(), m})
	w.Flush()
}
