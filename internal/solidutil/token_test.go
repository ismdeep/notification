package solidutil

import (
	"testing"
)

func TestRandAESKey(t *testing.T) {
	t.Logf("got = %v", RandAESKey())
}

func TestUserID(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				username: "admin",
			},
			want: "u_75Eatv6sZBHDWuSxc9GivXajgGHBuQBbAsUYLhV53Nkv",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UserID(tt.args.username); got != tt.want {
				t.Errorf("UserID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandToken(t *testing.T) {
	t.Logf("got = %v", RandToken())
}

func TestProjectID(t *testing.T) {
	type args struct {
		projectName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				projectName: "First Project",
			},
			want: "p_v9QgCHyphTzjsNCDvhTCXcjdG4gNyn9tNyBozyKwuGj",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProjectID(tt.args.projectName); got != tt.want {
				t.Errorf("ProjectID() = %v, want %v", got, tt.want)
			}
		})
	}
}
