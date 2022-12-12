package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/cosmos/cosmos-sdk/types/bech32"
)

func convert(addr string, target string) string {
	_, bz, err := bech32.DecodeAndConvert(addr)
	if err != nil {
		log.Fatal(err)
	}
	
	convertedAddr, err := bech32.ConvertAndEncode(target, bz)
	if err != nil {
		log.Fatal(err)
	}
	return convertedAddr
}

func main() {
	addr := os.Args[1]
	target := os.Args[2]

	f, err := os.Open(addr)
	if err != nil {
		convertedAddr := convert(addr, target)
		fmt.Printf("%s,%s\n", addr, convertedAddr)
	}
	
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		addr := scanner.Text()
		convertedAddr := convert(addr, target)
		fmt.Printf("%s,%s\n", addr, convertedAddr)
	}
}
