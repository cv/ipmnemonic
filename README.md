## IP Mnemonic

Generate a unique mnemonic for any IPv4 address.

Inspired by [@kseistrup](https://github.com/kseistrup)'s [Python implementation](https://github.com/kseistrup/ip-mnemonics) of [@gurno](https://github.com/gurno)'s writeup,
[IP Mnemonics](http://gurno.com/adam/mne/), here's a small Go utility that lets you convert IPv4 addresses to and from mnemonics:

```sh
$ ipmnemonic encode 10.0.0.1
arm-zip-zip-ace
$ ipmnemonic encode 127.0.0.1
lab-zip-zip-ace
$ ipmnemonic encode 10.0.0.1 172.16.2.3 192.168.4.255
arm-zip-zip-ace
pal-ban-act-add
rat-out-age-all
$ dig a www.google.com +short | xargs ipmnemonic encode
sky-eat-shy-hog
$ dig a times.com +short | xargs ipmnemonic encode
nab-hop-lap-odd
nab-hop-ace-odd
nab-hop-raw-odd
nab-hop-fan-odd
```

### Usage

```txt
usage: ipmnemonic [<flags>] <command> [<args> ...]

Encode and decode IP mnemonics

Flags:
      --help           Show context-sensitive help (also try --help-long and --help-man).
      --version        Show application version.
  -s, --separator="-"  Use char as separator

Commands:
  help [<command>...]
    Show help.

  encode <ip>...
    Encodes an IP address

  decode <mnemonic>...
    Decodes a mnemonic
```

### Installation

If you have a Go environment installed on your machine:
```
go get -u github.com/cv/ipmnemonic
```

Otherwise, grab a binary release from the [Releases](https://github.com/cv/ipmnemonic/releases) tab.

Enjoy! :smile:
