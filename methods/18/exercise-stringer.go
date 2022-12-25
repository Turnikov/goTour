package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func (p IPAddr) String() string {
	var output [4]string
	for i, value := range p {
		output[i] = strconv.FormatUint(uint64(value), 10)
	}
	return fmt.Sprint(strings.Join(output[:], "."))
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
