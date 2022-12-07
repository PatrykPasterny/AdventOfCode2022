package day_fifth

import "testing"

const (
	successfulTestFilepath   = "test_data/test_data.txt"
	nonExistingTestFilepath  = "test_data/not_existing_file.txt"
	noWhitespaceTestFilepath = "test_data/no_whitespace_between_inputs_test_data.txt"
)

func TestGetTopCratesNamesForCrateMover9000(t *testing.T) {
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
			want:    "CMZ",
			wantErr: false,
		},
		{
			name: "failing due to non existing input data",
			args: args{
				filepath: nonExistingTestFilepath,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "no whitespaces between crates and crane moves input test",
			args: args{
				filepath: noWhitespaceTestFilepath,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTopCratesNamesForCrateMover9000(tt.args.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTopCratesNamesForCrateMover9000() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetTopCratesNamesForCrateMover9000() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTopCratesNamesForCrateMover9001(t *testing.T) {
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
			want:    "MCD",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTopCratesNamesForCrateMover9001(tt.args.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTopCratesNamesForCrateMover9001() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetTopCratesNamesForCrateMover9001() got = %v, want %v", got, tt.want)
			}
		})
	}
}
