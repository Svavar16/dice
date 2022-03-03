package main

import "testing"

func Test_dice(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name   string
		args   args
		higher int
		lower  int
	}{
		{
			name:   "testing dice() with 5",
			args:   args{size: 5},
			higher: 5,
			lower:  0,
		},
		{
			name:   "testing dice() with 10",
			args:   args{size: 10},
			higher: 10,
			lower:  0,
		},
		{
			name:   "testing dice() with 15",
			args:   args{size: 15},
			higher: 15,
			lower:  0,
		},
		{
			name:   "testing dice() with 20",
			args:   args{size: 20},
			higher: 20,
			lower:  0,
		},
		{
			name:   "testing dice() with 25",
			args:   args{size: 25},
			higher: 25,
			lower:  0,
		},
		{
			name:   "testing dice() with 10",
			args:   args{size: 100},
			higher: 100,
			lower:  0,
		},
		{
			name:   "testing dice() with 200",
			args:   args{size: 200},
			higher: 200,
			lower:  0,
		},
		{
			name:   "testing dice() with 400",
			args:   args{size: 400},
			higher: 400,
			lower:  0,
		},
	}

	for _, tt := range tests {
		for i := 0; i < 50; i++ {
			t.Run(tt.name, func(t *testing.T) {
				if got := dice(tt.args.size); got <= tt.lower {
					t.Errorf("dice() = %v, lower or equal to %v", got, tt.lower)
				} else if got := dice(tt.args.size); got > tt.higher {
					t.Errorf("dice() = %v, higher than %v", got, tt.higher)
				}
			})
		}
	}
}
