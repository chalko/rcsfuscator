package obfuscator

import (
	"bytes"
	"regexp"
	"strings"
	"testing"
)

// TestHelloName calls obfuscator.Obfuscate with a name, checking
// for a valid return value.
func TestObfuscateName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	buf := new(bytes.Buffer)
	Obfuscate(strings.NewReader(name), buf)
	actual := buf.String()
	if !want.MatchString(actual) {
		t.Fatalf(`Obfuscate("Gladys") = %q, want match for %q, nil`, actual, want)
	}
}
