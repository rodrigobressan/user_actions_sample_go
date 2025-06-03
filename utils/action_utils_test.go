package utils

import (
	"reflect"
	"testing"
	"user_actions_sample_go/models"
)

func TestComputeNextActionProbsSorted(t *testing.T) {
	type args struct {
		actions    []models.Action
		actionType string
	}
	tests := []struct {
		name string
		args args
		want []ActionProbability
	}{
		{
			name: "No actions",
			args: args{
				actions:    []models.Action{},
				actionType: "VIEW_CONTACTS",
			},
			want: []ActionProbability{},
		},
		{
			name: "Multiple actions for same user properly computes the probabilities of the next action (3 ADD_TO_CRM, 1 REFER_USER)",
			args: args{
				actions: []models.Action{
					{UserID: 1, Type: "VIEW_CONTACTS"},
					{UserID: 1, Type: "ADD_TO_CRM"},
					{UserID: 1, Type: "VIEW_CONTACTS"},
					{UserID: 1, Type: "ADD_TO_CRM"},
					{UserID: 1, Type: "VIEW_CONTACTS"},
					{UserID: 1, Type: "ADD_TO_CRM"},
					{UserID: 1, Type: "VIEW_CONTACTS"},
					{UserID: 1, Type: "REFER_USER"},
					// Total: 3 ADD_TO_CRM, 1 REFER_USER (0.75, 0.25)
				},
				actionType: "VIEW_CONTACTS",
			},
			want: []ActionProbability{
				{Type: "ADD_TO_CRM", Probability: 0.75},
				{Type: "REFER_USER", Probability: 0.25},
			},
		},
		{
			name: "Multiple actions for different users properly computes the probabilities of the next action",
			args: args{
				actions: []models.Action{
					{UserID: 1, Type: "VIEW_CONTACTS"},
					{UserID: 1, Type: "ADD_TO_CRM"},
					{UserID: 1, Type: "VIEW_CONTACTS"},
					{UserID: 1, Type: "ADD_TO_CRM"},
					{UserID: 1, Type: "VIEW_CONTACTS"},
					{UserID: 1, Type: "REFER_USER"}, // single REFER_USER

					{UserID: 2, Type: "VIEW_CONTACTS"},
					{UserID: 2, Type: "ADD_TO_CRM"},
					{UserID: 2, Type: "VIEW_CONTACTS"},
					{UserID: 2, Type: "ADD_TO_CRM"},
				},
				actionType: "VIEW_CONTACTS",
			},
			want: []ActionProbability{
				{Type: "ADD_TO_CRM", Probability: 0.8},
				{Type: "REFER_USER", Probability: 0.2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// test for the optimized function
			if got := ComputeNextActionProbs(tt.args.actions, tt.args.actionType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ComputeNextActionProbs() = %v, want %v", got, tt.want)
			}

			// test for the original function (naive implementation)
			if got := ComputeNextActionProbsNaive(tt.args.actions, tt.args.actionType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ComputeNextActionProbsNaive() = %v, want %v", got, tt.want)
			}
		})
	}
}
