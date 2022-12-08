package day_seventh

import "testing"

const (
	successfulTestFilepath  = "test_data/test_data.txt"
	nonExistingTestFilepath = "test_data/not_existing_file.txt"
)

func TestSumAllDirectoriesAboveGivenSize(t *testing.T) {
	type args struct {
		filepath string
		size     int
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
				size:     100000,
			},
			want:    95437,
			wantErr: false,
		},
		{
			name: "failing due to non existing input data",
			args: args{
				filepath: nonExistingTestFilepath,
				size:     100000,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumAllDirectoriesAboveGivenSize(tt.args.filepath, tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("SumAllDirectoriesAboveGivenSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SumAllDirectoriesAboveGivenSize() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindSmallestDirSizeNeededToDeleteForSystemUpdate(t *testing.T) {
	type args struct {
		filepath   string
		memorySize int
		updateSize int
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
				filepath:   successfulTestFilepath,
				memorySize: 70000000,
				updateSize: 30000000,
			},
			want:    24933642,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindSmallestDirSizeNeededToDeleteForSystemUpdate(tt.args.filepath, tt.args.memorySize, tt.args.updateSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindSmallestDirSizeNeededToDeleteForSystemUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FindSmallestDirSizeNeededToDeleteForSystemUpdate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
