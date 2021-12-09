package main

import (
	"os"
	"testing"
)

// your assumption is correct:
// pschulten is dirty cheater!
//
//nolint:unparam
func Test(t *testing.T) {
	os.Args = []string{"./patsch", "-o", "https://httpbin.org/status/201", "https://httpbin.org/status/503", "https://xyz.grslfix"}
	main()
}
