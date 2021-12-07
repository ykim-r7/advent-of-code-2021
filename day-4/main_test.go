package main

import (
	"reflect"
	"testing"
)

func Test_removeIfContains(t *testing.T) {
	type args struct {
		slice []int
		this  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Empty",
			args: args{
				slice: []int{},
				this:  0,
			},
			want: []int{},
		},
		{
			name: "Remove Middle",
			args: args{
				slice: []int{1, 2, 3},
				this:  2,
			},
			want: []int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeIfContains(tt.args.slice, tt.args.this); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeIfContains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_selectNumber(t *testing.T) {
	type fields struct {
		sequences [][]int
	}
	type args struct {
		selectedNumber int
	}
	tests := []struct {
		name       string
		args       args
		fields     fields
		wantFields fields
	}{
		{
			name: "Empty",
			args: args{
				selectedNumber: 0,
			},
			fields: fields{
				sequences: [][]int{},
			},
			wantFields: fields{
				sequences: [][]int{},
			},
		},
		{
			name: "Number Not Selected",
			args: args{
				selectedNumber: 0,
			},
			fields: fields{
				sequences: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			wantFields: fields{
				sequences: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
		},
		{
			name: "Number Selected",
			args: args{
				selectedNumber: 3,
			},
			fields: fields{
				sequences: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			wantFields: fields{
				sequences: [][]int{
					{1, 2},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
		},
		{
			name: "Last Number In Sequence Selected",
			args: args{
				selectedNumber: 7,
			},
			fields: fields{
				sequences: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7},
				},
			},
			wantFields: fields{
				sequences: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{
				sequences: tt.fields.sequences,
			}
			b.selectNumber(tt.args.selectedNumber)
			if !reflect.DeepEqual(b.sequences, tt.wantFields.sequences) {
				t.Errorf("b.selectNumber(); b = %v, want %v", tt.fields, tt.wantFields)
			}
		})
	}
}

func TestBoard_hasBingo(t *testing.T) {
	type fields struct {
		sequences [][]int
	}
	tests := []struct {
		name      string
		fields    fields
		wantBingo bool
	}{
		{
			name: "No Bingo",
			fields: fields{
				sequences: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			wantBingo: false,
		},
		{
			name: "Bingo",
			fields: fields{
				sequences: [][]int{
					{1, 2, 3},
					{},
					{7, 8, 9},
				},
			},
			wantBingo: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{
				sequences: tt.fields.sequences,
			}
			if gotBingo := b.hasBingo(); gotBingo != tt.wantBingo {
				t.Errorf("Board.hasBingo() = %v, want %v", gotBingo, tt.wantBingo)
			}
		})
	}
}

func TestBoard_calculateScore(t *testing.T) {
	type fields struct {
		sequences [][]int
	}
	tests := []struct {
		name      string
		fields    fields
		wantScore int
	}{
		{
			name: "Empty",
			fields: fields{
				sequences: [][]int{
					{},
				},
			},
			wantScore: 0,
		},
		{
			name: "Bingo Scored",
			fields: fields{
				sequences: [][]int{
					{1, 2, 3},
					{},
					{7, 8, 9},
				},
			},
			wantScore: 30,
		},
		{
			name: "No Bingo Scored",
			fields: fields{
				sequences: [][]int{
					{1, 2, 3},
					{4},
					{7, 8, 9},
				},
			},
			wantScore: 34,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{
				sequences: tt.fields.sequences,
			}
			if gotScore := b.calculateScore(); gotScore != tt.wantScore {
				t.Errorf("Board.calculateScore() = %v, want %v", gotScore, tt.wantScore)
			}
		})
	}
}

func Test_createBoard(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name  string
		args  args
		wantB Board
	}{
		{
			name: "Empty",
			args: args{
				numbers: []int{},
			},
			wantB: Board{
				sequences: [][]int{},
			},
		},
		{
			name: "1x1 Board",
			args: args{
				numbers: []int{1},
			},
			wantB: Board{
				sequences: [][]int{{1}, {1}},
			},
		},
		{
			name: "5x5 Board",
			args: args{
				numbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25},
			},
			wantB: Board{
				sequences: [][]int{
					{1, 2, 3, 4, 5},
					{6, 7, 8, 9, 10},
					{11, 12, 13, 14, 15},
					{16, 17, 18, 19, 20},
					{21, 22, 23, 24, 25},
					{1, 6, 11, 16, 21},
					{2, 7, 12, 17, 22},
					{3, 8, 13, 18, 23},
					{4, 9, 14, 19, 24},
					{5, 10, 15, 20, 25},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotB := createBoard(tt.args.numbers); !reflect.DeepEqual(gotB, tt.wantB) {
				t.Errorf("createBoard() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func Test_winBingo(t *testing.T) {
	type args struct {
		inputFile string
	}
	tests := []struct {
		name      string
		args      args
		wantScore int
	}{
		{
			name: "example",
			args: args{
				"test-input",
			},
			wantScore: 4512,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotScore := winBingo(tt.args.inputFile); gotScore != tt.wantScore {
				t.Errorf("winBingo() = %v, want %v", gotScore, tt.wantScore)
			}
		})
	}
}

func Test_loseBingo(t *testing.T) {
	type args struct {
		inputFile string
	}
	tests := []struct {
		name      string
		args      args
		wantScore int
	}{
		{
			name: "example",
			args: args{
				"test-input",
			},
			wantScore: 1924,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotScore := loseBingo(tt.args.inputFile); gotScore != tt.wantScore {
				t.Errorf("loseBingo() = %v, want %v", gotScore, tt.wantScore)
			}
		})
	}
}
