package day_second

import "testing"

const (
	successfulTestFilepath = "test_data/test_data.txt"
	failingTestFilepath    = "not_existing_file.txt"
)

func TestCountStrategyGuideResult(t *testing.T) {
	type args struct {
		filepath string
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
			},
			want:    15,
			wantErr: false,
		},
		{
			name: "failing test",
			args: args{
				filepath: failingTestFilepath,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CountFirstStrategyGuideResult(tt.args.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountFirstStrategyGuideResult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CountFirstStrategyGuideResult() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountSecondStrategyGuideResult(t *testing.T) {
	type args struct {
		filepath string
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
			},
			want:    12,
			wantErr: false,
		},
		{
			name: "failing test",
			args: args{
				filepath: failingTestFilepath,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CountSecondStrategyGuideResult(tt.args.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountSecondStrategyGuideResult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CountSecondStrategyGuideResult() got = %v, want %v", got, tt.want)
			}
		})
	}
}
