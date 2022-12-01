package day_first

import "testing"

const (
	successfulTestFilepath = "test_data/test_data.txt"
	failingTestFilepath    = "test_data/failing_test_data.txt"
)

func TestGetBackpackWithMostCalories(t *testing.T) {
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
			want:    24000,
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
			got, err := GetBackpackWithMostCalories(tt.args.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBackpackWithMostCalories() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetBackpackWithMostCalories() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetThreeBackpacksWithMostCalories(t *testing.T) {
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
			want:    45000,
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
			got, err := GetThreeBackpacksWithMostCalories(tt.args.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetThreeBackpacksWithMostCalories() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetThreeBackpacksWithMostCalories() got = %v, want %v", got, tt.want)
			}
		})
	}
}
