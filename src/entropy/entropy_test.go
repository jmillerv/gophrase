package entropy

import (
	"reflect"
	"testing"
)

func TestGetEntropy(t *testing.T) {
	type args struct {
		passphrase string
	}
	tests := []struct {
		name string
		args args
		want *Metadata
	}{
		{
			name: "Success: Returns Expected Metadata",
			args: args{
				passphrase: "craveprizemoist",
			},
			want: &Metadata{
				Entropy:   36.21,
				CrackTime: "3.0 months",
				Score:     3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEntropy(tt.args.passphrase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEntropy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintEntropy(t *testing.T) {
	type args struct {
		passphrase string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				passphrase: "vaNquished68eMperor9squirrelGetawaY/rustproof284564+",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintEntropy(tt.args.passphrase)
		})
	}
}
