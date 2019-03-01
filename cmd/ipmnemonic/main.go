package main

import (
	"fmt"
	"os"

	"github.com/cv/ipmnemonic/pkg/ipmnemonic"

	"github.com/alecthomas/kingpin"
)

var (
	app = kingpin.
		New("ipmnemonic", "Encode and decode IP mnemonics").
		Version(ipmnemonic.Version + " (" + ipmnemonic.Revision + ")")

	encode = app.Command("encode", "Encodes an IP address")
	ips    = encode.Arg("ip", "IPv4 addresses to encode").Required().IPList()

	decode = app.Command("decode", "Decodes a mnemonic")
	strs   = decode.Arg("mnemonic", "mnemonics to decode").Required().Strings()

	sep = app.Flag("separator", "Use char as separator").
		Short('s').
		Default("-").
		String()
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {

	case encode.FullCommand():
		for _, ip := range *ips {
			fmt.Println(ipmnemonic.Encode(ip.To4(), *sep))
		}

	case decode.FullCommand():
		for _, str := range *strs {
			fmt.Println(ipmnemonic.Decode(str, *sep))
		}
	}
}
