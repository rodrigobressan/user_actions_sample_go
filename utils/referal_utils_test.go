package utils

import (
	"reflect"
	"surfe_assignment/models"
	"testing"
)

func TestComputeReferralIndex(t *testing.T) {
	type args struct {
		actions []models.Action
	}
	tests := []struct {
		name string
		args args
		want map[int]int
	}{
		{
			name: "No actions",
			args: args{
				actions: []models.Action{},
			},
			want: map[int]int{},
		},
		{
			name: "Single user with no referrals",
			args: args{
				actions: []models.Action{
					{UserID: 1, Type: "VIEW_CONTACTS"},
				},
			},
			want: map[int]int{1: 0},
		},
		{
			name: "Single user with one referral",
			args: args{
				actions: []models.Action{
					{UserID: 1, Type: "REFER_USER", TargetUser: makeIntPointer(2)},
				},
			},
			want: map[int]int{1: 1, 2: 0}, // User 1 referred User 2, User 2 has no referrals
		},
		{
			name: "Multiple users with referrals and no referrals",
			args: args{
				actions: []models.Action{
					//User 1
					//├── User 2
					//└── User 3
					//    └── User 4
					//        └── User 5

					{UserID: 1, Type: "REFER_USER", TargetUser: makeIntPointer(2)},
					{UserID: 1, Type: "REFER_USER", TargetUser: makeIntPointer(3)},
					{UserID: 3, Type: "REFER_USER", TargetUser: makeIntPointer(4)},
					{UserID: 4, Type: "REFER_USER", TargetUser: makeIntPointer(5)},
				},
			},
			want: map[int]int{
				1: 4, // User 1 referred Users 2 and 3 (directly) and 4 and 5 (indirectly)
				2: 0, // User 2 has no referrals
				3: 2, // User 3 referred User 4 (directly) and User 5 (indirectly)
				4: 1, // User 4 referred User 5 (directly)
				5: 0, // User 5 has no referrals
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeReferralIndex(tt.args.actions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ComputeReferralIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func makeIntPointer(i int) *int {
	p := new(int)
	*p = i
	return p
}
