package roman

import (
	"testing"
)

func TestFromInt(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "1",
			args:    args{number: 1},
			want:    "I",
			wantErr: false,
		},
		{
			name:    "2",
			args:    args{number: 2},
			want:    "II",
			wantErr: false,
		},
		{
			name:    "3",
			args:    args{number: 3},
			want:    "III",
			wantErr: false,
		},
		{
			name:    "4",
			args:    args{number: 4},
			want:    "IV",
			wantErr: false,
		},
		{
			name:    "5",
			args:    args{number: 5},
			want:    "V",
			wantErr: false,
		},
		{
			name:    "6",
			args:    args{number: 6},
			want:    "VI",
			wantErr: false,
		},
		{
			name:    "7",
			args:    args{number: 7},
			want:    "VII",
			wantErr: false,
		},
		{
			name:    "8",
			args:    args{number: 8},
			want:    "VIII",
			wantErr: false,
		},
		{
			name:    "9",
			args:    args{number: 9},
			want:    "IX",
			wantErr: false,
		},
		{
			name:    "10",
			args:    args{number: 10},
			want:    "X",
			wantErr: false,
		},
		{
			name:    "error 0",
			args:    args{number: 0},
			want:    "",
			wantErr: true,
		},
		{
			name:    "error more then 10",
			args:    args{number: 12},
			want:    "",
			wantErr: true,
		},
		{
			name:    "error less then 0",
			args:    args{number: -2},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromInt(tt.args.number)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FromInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromRoman(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "I",
			args:    args{s: "I"},
			want:    1,
			wantErr: false,
		},
		{
			name:    "II",
			args:    args{s: "II"},
			want:    2,
			wantErr: false,
		},
		{
			name:    "III",
			args:    args{s: "III"},
			want:    3,
			wantErr: false,
		},
		{
			name:    "IV",
			args:    args{s: "IV"},
			want:    4,
			wantErr: false,
		},
		{
			name:    "V",
			args:    args{s: "V"},
			want:    5,
			wantErr: false,
		},
		{
			name:    "VI",
			args:    args{s: "VI"},
			want:    6,
			wantErr: false,
		},
		{
			name:    "VII",
			args:    args{s: "VII"},
			want:    7,
			wantErr: false,
		},
		{
			name:    "VIII",
			args:    args{s: "VIII"},
			want:    8,
			wantErr: false,
		},
		{
			name:    "IX",
			args:    args{s: "IX"},
			want:    9,
			wantErr: false,
		},
		{
			name:    "X",
			args:    args{s: "X"},
			want:    10,
			wantErr: false,
		},
		{
			name:    "more then 10, one byte",
			args:    args{s: "D"},
			want:    0,
			wantErr: true,
		},
		{
			name:    "more then 10, two bytes",
			args:    args{s: "XI"},
			want:    0,
			wantErr: true,
		},
		{
			name:    "unknown roman number, one byte",
			args:    args{s: "K"},
			want:    0,
			wantErr: true,
		},
		{
			name:    "unknown roman number, last byte",
			args:    args{s: "XK"},
			want:    0,
			wantErr: true,
		},
		{
			name:    "unknown roman number, first byte",
			args:    args{s: "KX"},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromRoman(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromRoman() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FromRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
