package main

import (
	"reflect"
	"testing"
)

func Test_sumColumns(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name     string
		args     args
		wantSums []int
	}{
		{
			name: "example",
			args: args{[]string{
				"00100",
				"11110",
				"10110",
				"10111",
				"10101",
				"01111",
				"00111",
				"11100",
				"10000",
				"11001",
				"00010",
				"01010",
			}},
			wantSums: []int{7, 5, 8, 7, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSums := sumColumns(tt.args.lines); !reflect.DeepEqual(gotSums, tt.wantSums) {
				t.Errorf("sumColumns() = %v, want %v", gotSums, tt.wantSums)
			}
		})
	}
}

func Test_getCo2(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name    string
		args    args
		wantCo2 uint64
	}{
		{
			name: "example",
			args: args{[]string{
				"00100",
				"11110",
				"10110",
				"10111",
				"10101",
				"01111",
				"00111",
				"11100",
				"10000",
				"11001",
				"00010",
				"01010",
			}},
			wantCo2: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCo2 := getCo2(tt.args.lines); gotCo2 != tt.wantCo2 {
				t.Errorf("getCo2() = %v, want %v", gotCo2, tt.wantCo2)
			}
		})
	}
}

func Test_getO2(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name   string
		args   args
		wantO2 uint64
	}{
		{
			name: "example",
			args: args{[]string{
				"00100",
				"11110",
				"10110",
				"10111",
				"10101",
				"01111",
				"00111",
				"11100",
				"10000",
				"11001",
				"00010",
				"01010",
			}},
			wantO2: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotO2 := getO2(tt.args.lines); gotO2 != tt.wantO2 {
				t.Errorf("getO2() = %v, want %v", gotO2, tt.wantO2)
			}
		})
	}
}
