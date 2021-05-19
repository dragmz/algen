package algen

import (
	"runtime"
	"strings"

	"github.com/algorand/go-algorand-sdk/crypto"
)

type GenerateArgs struct {
	StartsWith string
	EndsWith   string
}

func accept(a crypto.Account, args GenerateArgs) bool {
	if args.StartsWith != "" {
		if !strings.HasPrefix(a.Address.String(), args.StartsWith) {
			return false
		}
	}

	if args.EndsWith != "" {
		if !strings.HasSuffix(a.Address.String(), args.EndsWith) {
			return false
		}
	}

	return true
}

func GenerateAddress(args GenerateArgs) (crypto.Account, error) {
	done := make(chan struct{})
	ch := make(chan crypto.Account)

	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for {
				a := crypto.GenerateAccount()

				if accept(a, args) {
					select {
					case ch <- a:
					case <-done:
					}

					return
				}
			}
		}()
	}

	a := <-ch
	close(done)

	return a, nil
}
