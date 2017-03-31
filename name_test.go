package ipisp

import (
	"reflect"
	"testing"
)

func TestParseName(t *testing.T) {
	type args struct {
		raw string
	}
	tests := []struct {
		name string
		args args
		want Name
	}{
		{"no long", args{"google"}, Name{Raw: "google", Short: "google", Long: "google"}},
		{"long", args{"google - GOOGLE stuff"}, Name{Raw: "google - GOOGLE stuff", Short: "google", Long: "GOOGLE stuff"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseName(tt.args.raw); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseName() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
