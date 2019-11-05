package hello

import "testing"

func TestHi(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"Hi!"}, "Hello World!\nHi!\n"},
		{"2", args{"Fails!"}, "Hello World!\nHi!\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hi(tt.args.message); got != tt.want {
				t.Errorf("Hi() = %v\nwant %v", got, tt.want)
			}
		})
	}
}
