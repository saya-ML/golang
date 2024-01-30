package main

import "testing"

func TestRun(t *testing.T) {
	type args struct {
		workersCount int
		worksToDo    int
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "all empty",
			args: args{
				workersCount: 0,
				worksToDo:    0,
			},
			want: 0,
		},
		{
			name: "500 - 1_000",
			args: args{
				workersCount: 500,
				worksToDo:    1_000,
			},
			want: 1_000,
		},
		{
			name: "1_001 - 1_000",
			args: args{
				workersCount: 1_001,
				worksToDo:    1_000,
			},
			want: 1_000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Run(tt.args.workersCount, tt.args.worksToDo); got != tt.want {
				t.Errorf("Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
