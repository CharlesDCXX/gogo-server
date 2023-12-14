package main

import (
	"context"
	"testing"
)

func TestGetChanNosAndConsistency1(t *testing.T) {
	type args struct {
		context        context.Context
		ProductFeature string
	}
	tests := []struct {
		name            string
		args            args
		wantChanNos     string
		wantConsistency int
		wantErr         bool
	}{
		// TODO: Add test cases.
		{"name1", args{context: context.TODO(), ProductFeature: "{\"feature\": [4107,4112],\"chan_nos\":\"[1,2,3]\",\"chan_consistency\":1}"}, "[1,2,3]", 1, false},
		// {"name2", args{context: context.TODO(), ProductFeature: ""}, "1", 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotChanNos, gotConsistency, err := GetChanNosAndConsistency(tt.args.context, tt.args.ProductFeature)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetChanNosAndConsistency() error = %v, wantErr %v", err, tt.wantErr)

			}
			if gotChanNos != tt.wantChanNos {
				t.Errorf("GetChanNosAndConsistency() gotChanNos = %v, want %v", gotChanNos, tt.wantChanNos)
			}
			if gotConsistency != tt.wantConsistency {
				t.Errorf("GetChanNosAndConsistency() gotConsistency = %v, want %v", gotConsistency, tt.wantConsistency)
			}
		})
	}
}
