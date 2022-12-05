package day_third

import "testing"

const (
	successfulTestFilepath                        = "test_data/test_data.txt"
	nonExistingTestFilepath                       = "test_data/not_existing_file.txt"
	lackOfCommonLetterTestFilepath                = "test_data/first_failing_test.txt"
	commonLettersOutOfEnglishAlphabetTestFilepath = "test_data/second_failing_test.txt"
	notEvenRucksacksLengthTestFilepath            = "test_data/third_failing_test.txt"
)

func TestSumRucksacksPrioritySupplies(t *testing.T) {
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
			want:    157,
			wantErr: false,
		},
		{
			name: "failing due to input not containing common letters in same rucksack's components",
			args: args{
				filepath: lackOfCommonLetterTestFilepath,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "failing due to having common letters out of english alphabet within rucksacks components",
			args: args{
				filepath: commonLettersOutOfEnglishAlphabetTestFilepath,
			},
			want:    0,
			wantErr: true,
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
			name: "failing due to rucksacks' length not being even",
			args: args{
				filepath: notEvenRucksacksLengthTestFilepath,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumRucksacksPrioritySupplies(tt.args.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("SumRucksacksPrioritySupplies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SumRucksacksPrioritySupplies() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumGroupBadgesPriorities(t *testing.T) {
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
			want:    70,
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumGroupBadgesPriorities(tt.args.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("SumGroupBadgesPriorities() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SumGroupBadgesPriorities() got = %v, want %v", got, tt.want)
			}
		})
	}
}
