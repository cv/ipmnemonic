package ipmnemonic

import (
	"bytes"
	"net"
)

func Encode(ip net.IP, sep string) string {
	out := bytes.NewBuffer([]byte{})
	for i, b := range ip {
		if i != 0 {
			out.WriteString(sep)
		}
		out.WriteString(Words[b])
	}

	return out.String()
}
