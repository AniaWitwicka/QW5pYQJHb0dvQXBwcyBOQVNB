package operations

import (
	"reflect"
	"testing"
	"time"
)

func getParsedDate(s string) time.Time {
	date, _ := time.Parse("2006-01-02", s)
	return date
}

func TestGetUrlLen(t *testing.T) {
	type args struct {
		from time.Time
		to   time.Time
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Correct",
			args: args{
				from: getParsedDate("2022-07-01"),
				to:   getParsedDate("2022-07-01"),
			},
			want: 1,
		},
		{
			name: "Test Correct",
			args: args{
				from: getParsedDate("2022-07-01"),
				to:   getParsedDate("2022-07-02"),
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUrlLen(tt.args.from, tt.args.to); got != tt.want {
				t.Errorf("GetUrlLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseAndValidateDateRanges(t *testing.T) {
	type args struct {
		from string
		to   string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		want1   time.Time
		wantErr bool
	}{
		{
			name : "Valid date ranges",
			args: args{
				from: "2022-07-01",
				to:   "2022-07-01",
			},
			want: getParsedDate("2022-07-01"),
			want1: getParsedDate("2022-07-01"),
			wantErr: false,
		},
		{
			name : "Valid date ranges",
			args: args{
				from: "2022-07-01",
				to:   "2022-07-02",
			},
			want: getParsedDate("2022-07-01"),
			want1: getParsedDate("2022-07-02"),
			wantErr: false,
		},
		{
			name : "From after To Date",
			args: args{
				from: "2022-07-04",
				to:   "2022-07-02",
			},
			want: time.Time{},
			want1: time.Time{},
			wantErr: true,
		},
		{
			name : "Invalid From date",
			args: args{
				from: "invalid Date",
				to:   "2022-07-02",
			},
			want: time.Time{},
			want1: time.Time{},
			wantErr: true,
		},
		{
			name : "Invalid To date",
			args: args{
				from: "2022-07-02",
				to:   "invalid Date",
			},
			want: time.Time{},
			want1: time.Time{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := ParseAndValidateDateRanges(tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseAndValidateDateRanges() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseAndValidateDateRanges() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ParseAndValidateDateRanges() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}