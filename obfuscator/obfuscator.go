package obfuscator

import (
	"bufio"
	"inet.af/netaddr"
	"io"
)

var TEST_NET_1 = netaddr.MustParseIPPrefix("192.0.2.0/24")

//RCS Obuscation config
type Rob struct {
	addresses []string
	vlans     []string
}

func NewRob(addresses []string, vlans []string) *Rob {
	r := Rob{addresses: addresses, vlans: vlans}

	return &r
}

func Obfuscate(in io.Reader, out io.Writer) {
	// read the file line by line using scanner
	scanner := bufio.NewScanner(in)
	buffout := bufio.NewWriter(out)
	for scanner.Scan() {
		line := scanner.Text()
		buffout.WriteString(obfuscate(line))
		buffout.Flush()
	}
}

func obfuscate(orig string) string {
	return orig
}
