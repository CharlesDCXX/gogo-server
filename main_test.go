package main

import (
	"context"
	"fmt"
	"testing"
)

func TestGetChanNosAndConsistency(t *testing.T) {
	type args struct {
		context        context.Context
		ProductFeature string
	}
	tests := []struct {
		name       string
		args       args
		wantResult map[string]int
		wantErr    bool
	}{
		// TODO: Add test cases.
		{"name1", args{context: context.TODO(), ProductFeature: "{\"feature\": [4107,4112],\"chan_nos\":\"[1,2,3]\",\"cloud_info_chan_map\": {\"1\": 1,\"2\": 2}}"}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := GetChanNosAndConsistency(tt.args.context, tt.args.ProductFeature)
			if err != nil {
				fmt.Printf("test err: %v\n", err)
			}
			for k, v := range gotResult {
				fmt.Printf("key: %v,value: %v\n", k, v)
			}
		})
	}
}
