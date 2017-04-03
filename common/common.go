package common

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BytesToString(bs []byte) string {
	l := len(bs)
	buf := make([]string, 0, l)
	for i := 0; i < l; i++ {
		buf = appendString(buf, bs[i])
	}
	return strings.Join(buf, ".")
}

func appendString(bs []string, b byte) []string {
	var a byte
	var s int
	for i := 0; i < 8; i++ {
		a = b
		b <<= 1
		b >>= 1
		switch a {
		case b:
			s += 0
		default:
			temp := 1
			for j := 0; j < 7-i; j++ {
				temp = temp * 2
			}
			s += temp
		}

		b <<= 1
	}

	return append(bs, strconv.Itoa(s))
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
func Log(v ...interface{}) {
	fmt.Println(v...)
}
