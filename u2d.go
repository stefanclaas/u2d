
package main

import (
	"bufio"
	"flag"
	"fmt"
	"strings"
	"math/big"
	"os"
)

func main() {
	decode := flag.Bool("d", false, "Decode a decimal string to text")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 0, 64*1024), 10*1024*1024)

	for scanner.Scan() {
		line := scanner.Text()

		if *decode {
			for i := 0; i < len(line); i += 7 {
				number := line[i : i+7]
				bigInt := new(big.Int)
				bigInt.SetString(strings.TrimSpace(number), 10)
				i := bigInt.Int64()
				fmt.Print(string(i))
			}
			fmt.Print("\n")
		} else {
			for _, r := range line {
				bigInt := new(big.Int)
				bigInt.SetInt64(int64(r))
				fmt.Printf("%07s", bigInt.String())
			}
			//bigInt := new(big.Int)
			//bigInt.SetInt64(int64('\n'))
			//fmt.Printf("%07s", bigInt.String())
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
	}

	fmt.Println()
}

