package day_tenth

import (
	"reflect"
	"strings"
	"testing"
)

const (
	successfulTestFilepath       = "test_data/test_data.txt"
	successfulScreenTestFilepath = "test_data/screen_test_data.txt"
	nonExistingTestFilepath      = "test_data/not_existing_file.txt"
	wrongAddxTypeTestFilepath    = "test_data/addx_not_int_test_data.txt"
)

func TestGetSumOfStrengthsOfGivenCycles(t *testing.T) {
	testCycles := []int{20, 60, 100, 140, 180, 220}

	type args struct {
		filepath string
		cycles   []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "successful test",
			args: args{
				filepath: successfulTestFilepath,
				cycles:   testCycles,
			},
			want:    13140,
			wantErr: false,
		},
		{
			name: "failing due to non existing input file",
			args: args{
				filepath: nonExistingTestFilepath,
				cycles:   testCycles,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "failing due to wrongly typed X register increment near addx command data",
			args: args{
				filepath: wrongAddxTypeTestFilepath,
				cycles:   testCycles,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSumOfStrengthsOfGivenCycles(tt.args.filepath, tt.args.cycles)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSumOfStrengthsOfGivenCycles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetSumOfStrengthsOfGivenCycles() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetScreenOutput(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "successful test",
			args: args{
				filepath: successfulTestFilepath,
			},
			want:    "##..##..##..##..##..##..##..##..##..##..###...###...###...###...###...###...###.####....####....####....####....####....#####.....#####.....#####.....#####.....######......######......######......###########.......#######.......#######.....",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetScreenOutput(tt.args.filepath)
			sb := strings.Builder{}
			for i := range got {
				for j := range got[i] {
					sb.WriteString(got[i][j])
				}
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("GetScreenOutput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			gotString := sb.String()

			if !reflect.DeepEqual(gotString, tt.want) {
				t.Errorf("GetScreenOutput() got = %v, want %v", gotString, tt.want)
			}
		})
	}
}
