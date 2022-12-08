package day_sixth

import "testing"

const (
	firstSuccessfulTestFilepath  = "test_data/first_test_data.txt"
	secondSuccessfulTestFilepath = "test_data/sec_test_data.txt"
	thirdSuccessfulTestFilepath  = "test_data/third_test_data.txt"
	fourthSuccessfulTestFilepath = "test_data/fourth_test_data.txt"
	fifthSuccessfulTestFilepath  = "test_data/fifth_test_data.txt"

	nonExistingTestFilepath    = "test_data/not_existing_file.txt"
	tooShortSignalTestFilepath = "test_data/too_short_signal_test_data.txt"
	tooManyLinesTestFilepath   = "test_data/too_many_lines_test_data.txt"
	noMarkerFoundTestFilepath  = "test_data/no_marker_found_test_data.txt"
)

func TestFindFirstStartOfPacketMarker(t *testing.T) {
	type args struct {
		filepaths []string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr []bool
	}{
		{
			name: "successful test",
			args: args{
				filepaths: []string{
					firstSuccessfulTestFilepath,
					secondSuccessfulTestFilepath,
					thirdSuccessfulTestFilepath,
					fourthSuccessfulTestFilepath,
					fifthSuccessfulTestFilepath,
				},
			},
			want:    []int{7, 5, 6, 10, 11},
			wantErr: []bool{false, false, false, false, false},
		},
		{
			name: "failing due to non existing input data",
			args: args{
				filepaths: []string{nonExistingTestFilepath},
			},
			want:    []int{0},
			wantErr: []bool{true},
		},
		{
			name: "failing signal being too short to find a marker",
			args: args{
				filepaths: []string{tooShortSignalTestFilepath},
			},
			want:    []int{0},
			wantErr: []bool{true},
		},
		{
			name: "failing due to input having too many lines",
			args: args{
				filepaths: []string{tooManyLinesTestFilepath},
			},
			want:    []int{0},
			wantErr: []bool{true},
		},
		{
			name: "failing due to lack of packet starting marker",
			args: args{
				filepaths: []string{noMarkerFoundTestFilepath},
			},
			want:    []int{0},
			wantErr: []bool{true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for idx := range tt.args.filepaths {
				got, err := FindFirstStartOfPacketMarker(tt.args.filepaths[idx])
				if (err != nil) != tt.wantErr[idx] {
					t.Errorf("FindFirstStartOfPacketMarker() error = %v, wantErr %v", err, tt.wantErr[idx])
					return
				}
				if got != tt.want[idx] {
					t.Errorf("FindFirstStartOfPacketMarker() got = %v, want %v", got, tt.want[idx])
				}
			}
		})
	}
}

func TestFindFirstStartOfMessageMarker(t *testing.T) {
	type args struct {
		filepaths []string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr []bool
	}{
		{
			name: "successful test",
			args: args{
				filepaths: []string{
					firstSuccessfulTestFilepath,
					secondSuccessfulTestFilepath,
					thirdSuccessfulTestFilepath,
					fourthSuccessfulTestFilepath,
					fifthSuccessfulTestFilepath,
				},
			},
			want:    []int{19, 23, 23, 29, 26},
			wantErr: []bool{false, false, false, false, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for idx := range tt.args.filepaths {
				got, err := FindFirstStartOfMessageMarker(tt.args.filepaths[idx])
				if (err != nil) != tt.wantErr[idx] {
					t.Errorf("FindFirstStartOfMessageMarker() error = %v, wantErr %v", err, tt.wantErr[idx])
					return
				}
				if got != tt.want[idx] {
					t.Errorf("FindFirstStartOfMessageMarker() got = %v, want %v", got, tt.want[idx])
				}
			}
		})
	}
}
