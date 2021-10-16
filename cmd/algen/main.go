package main

import (
	"encoding/csv"
	"flag"
	"os"
	"strings"

	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/mnemonic"
	"github.com/dragmz/algen"
)

func printAccount(a crypto.Account) {
	m, err := mnemonic.FromPrivateKey(a.PrivateKey)
	if err != nil {
		panic(err)
	}

	w := csv.NewWriter(os.Stdout)
	w.Write([]string{a.Address.String(), m})
	w.Flush()
}

func main() {
	var startsWith string
	var endsWith string
	var count uint64
	var message string
	var chunk int
	var contains string

	flag.StringVar(&startsWith, "starts-with", "", "generate address that starts with a given string")
	flag.StringVar(&endsWith, "ends-with", "", "generate address that ends with a given string")
	flag.StringVar(&contains, "contains", "", "generate address that contains a given string")
	flag.StringVar(&message, "message", "", "generate addresses with prefixes to compose given message")
	flag.Uint64Var(&count, "count", 1, "number of addresses to generate")
	flag.IntVar(&chunk, "chunk", 3, "chunk length for encoding message")
	flag.Parse()

	if message != "" {
		message = strings.ReplaceAll(strings.ToUpper(message), " ", "")
		var chunks []string

		parts := len(message) / chunk
		if parts*chunk < len(message) {
			parts += 1
		}

		for i := 0; i < parts; i++ {
			to := i*chunk + chunk
			if to > len(message) {
				to = len(message)
			}
			chunks = append(chunks, message[i*chunk:to])
		}

		for _, item := range chunks {
			a, err := algen.GenerateAddress(algen.GenerateArgs{
				StartsWith: item,
			})
			if err != nil {
				panic(err)
			}

			printAccount(a)
		}
	} else {
		for i := uint64(0); i < count; i++ {
			a, err := algen.GenerateAddress(algen.GenerateArgs{
				StartsWith: strings.ToUpper(startsWith),
				EndsWith:   strings.ToUpper(endsWith),
				Contains:   strings.Split(strings.ToUpper(contains), ","),
			})

			if err != nil {
				panic(err)
			}

			printAccount(a)
		}
	}
}
