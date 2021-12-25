package ipisp

import "testing"

func TestParseASN(t *testing.T) {
	type args struct {
		asn string
	}
	tests := []struct {
		name    string
		args    args
		want    ASN
		wantErr bool
	}{
		{"", args{"AS555"}, ASN(555), false},
		{"", args{"AS"}, ASN(-1), true},
		{"", args{"ASDFASDF"}, ASN(-1), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseASN(tt.args.asn)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseASN() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseASN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestASN_String(t *testing.T) {
	tests := []struct {
		name string
		a    ASN
		want string
	}{
		{"", ASN(555), "AS555"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.String(); got != tt.want {
				t.Errorf("ASN.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
