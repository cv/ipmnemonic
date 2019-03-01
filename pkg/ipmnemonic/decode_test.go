package ipmnemonic

import (
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	tests := []struct {
		str  string
		sep  string
		want string
	}{
		{"zip-zip-zip-zip", "-", "0.0.0.0"},
		{"arm-zip-zip-bet", "-", "10.0.0.21"},
		{"all-all-all-all", "-", "255.255.255.255"},
	}
	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			if have := Decode(tt.str, tt.sep); !reflect.DeepEqual(have, tt.want) {
				t.Errorf("Decode() = %v, want %v", have, tt.want)
			}
		})
	}
}
