package day_fourth

import "testing"

const (
	successfulTestFilepath           = "test_data/test_data.txt"
	nonExistingTestFilepath          = "test_data/not_existing_file.txt"
	tooBigGroupsTestFilepath         = "test_data/too_big_group_test_data.txt"
	tooManyAssignmentsTestFilepath   = "test_data/too_many_assignments_test_data.txt"
	assignmentNotANumberTestFilepath = "test_data/assignment_not_a_number_test_data.txt"
)

func TestSumFullyOverlappingAssignments(t *testing.T) {
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
			want:    2,
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
			name: "failing due to input containing to big groups",
			args: args{
				filepath: tooBigGroupsTestFilepath,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "failing due to elves having too many assignments",
			args: args{
				filepath: tooManyAssignmentsTestFilepath,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "failing due to assignments not being numbers",
			args: args{
				filepath: assignmentNotANumberTestFilepath,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumFullyOverlappingAssignments(tt.args.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("SumFullyOverlappingAssignments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SumFullyOverlappingAssignments() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumEveryOverlappingAssignments(t *testing.T) {
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
			want:    4,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumEveryOverlappingAssignments(tt.args.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("SumEveryOverlappingAssignments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SumEveryOverlappingAssignments() got = %v, want %v", got, tt.want)
			}
		})
	}
}
