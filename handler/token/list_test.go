package token

import (
	"testing"
)

func TestList(t *testing.T) {
	type args struct {
		userID uint
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestList-001",
			args: args{
				userID: 1000,
			},
			wantErr: true,
		},
		{
			name: "TestList-002",
			args: args{
				userID: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := List(tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
