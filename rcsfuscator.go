package main

import (
	"bufio"
	"chalko.com/obfuscator"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	obfuscator.Obfuscate(reader, writer)
}
