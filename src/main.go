package main

import (
	"fmt"

	"github.com/themobilecoder/ocm-meta/meta"
)

func main() {
	monkeys := meta.GetOnChainMonkeys()
	fmt.Print(monkeys[4642-1])
}
