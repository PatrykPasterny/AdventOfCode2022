package day_eighth

import "testing"

const (
	successfulTestFilepath         = "test_data/test_data.txt"
	nonExistingTestFilepath        = "test_data/not_existing_file.txt"
	nonNumericalMatrixTestFilepath = "test_data/non_numerical_matrix_elements_test_data.txt"
)

func TestCountVisibleTrees(t *testing.T) {
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
			want:    21,
			wantErr: false,
		},
		{
			name: "failing due to non existing input data",
			args: args{
				filepath: nonExistingTestFilepath,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "failing due to non numerical input data within tree grid",
			args: args{
				filepath: nonNumericalMatrixTestFilepath,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CountVisibleTrees(tt.args.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountVisibleTrees() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CountVisibleTrees() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindHighestScenicScore(t *testing.T) {
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
			want:    8,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindHighestScenicScore(tt.args.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindHighestScenicScore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FindHighestScenicScore() got = %v, want %v", got, tt.want)
			}
		})
	}
}
