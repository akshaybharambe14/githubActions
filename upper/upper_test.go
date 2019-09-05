package upper

import "testing"

func TestToUpper(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"valid",
			args{s: "this should be in upper case"},
			"THIS SHOULD BE IN UPPER CASE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUpper(tt.args.s); got != tt.want {
				t.Errorf("ToUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}
