package ipmnemonic

import (
	"net"
	"reflect"
	"testing"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		str  string
		sep  string
		want string
	}{
		{"0.0.0.0", "-", "zip-zip-zip-zip"},
		{"10.0.0.21", "-", "arm-zip-zip-bet"},
		{"255.255.255.255", "-", "all-all-all-all"},
	}
	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			if have := Encode(net.ParseIP(tt.str).To4(), tt.sep); !reflect.DeepEqual(have, tt.want) {
				t.Errorf("Decode() = %v, want %v", have, tt.want)
			}
		})
	}
}
