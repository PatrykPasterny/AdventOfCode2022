package day_ninth

import (
	"testing"
)

const (
	successfulTestFilepath     = "test_data/test_data.txt"
	successfulBigTestFilepath  = "test_data/big_test_data.txt"
	nonExistingTestFilepath    = "test_data/not_existing_file.txt"
	wrongStepsTypeTestFilepath = "wrong_steps_type_test_data.txt"
	undefinedMoveTestFilepath  = "undefined_move_test_data.txt"
)

func TestCountAllTailPositions1(t *testing.T) {
	type args struct {
		filepath   string
		ropeLength int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "successful 2 nodes rope test",
			args: args{
				filepath:   successfulTestFilepath,
				ropeLength: 2,
			},
			want:    13,
			wantErr: false,
		},
		{
			name: "successful 10 nodes rope test",
			args: args{
				filepath:   successfulTestFilepath,
				ropeLength: 10,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "successful 10 nodes rope test",
			args: args{
				filepath:   successfulBigTestFilepath,
				ropeLength: 10,
			},
			want:    36,
			wantErr: false,
		},
		{
			name: "failing due to rope being to short input data",
			args: args{
				filepath:   successfulTestFilepath,
				ropeLength: 1,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "failing due to non existing input data",
			args: args{
				filepath:   nonExistingTestFilepath,
				ropeLength: 2,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "failing due to steps not being int in input data",
			args: args{
				filepath:   wrongStepsTypeTestFilepath,
				ropeLength: 2,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "failing due to direction not being U, R, D or L in input data",
			args: args{
				filepath:   undefinedMoveTestFilepath,
				ropeLength: 2,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CountAllTailPositions(tt.args.filepath, tt.args.ropeLength)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountAllTailPositions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CountAllTailPositions() got = %v, want %v", got, tt.want)
			}
		})
	}
}
