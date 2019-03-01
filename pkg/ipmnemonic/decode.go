package ipmnemonic

import (
	"net"
	"strings"
)

func Decode(s, sep string) string {
	var out []byte
	for _, iw := range strings.Split(s, sep) {
		for j, jw := range Words {
			if iw == jw {
				out = append(out, byte(j))
				break
			}
		}
	}
	return net.IP(out).To4().String()
}
