package token

import (
	"github.com/ismdeep/notification/api/request"
	"testing"
)

func TestNewToken(t *testing.T) {
	type args struct {
		userID uint
		req    *request.NewToken
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "TestNewToken-001",
			args:    args{
				userID: 0,
				req:    nil,
			},
			wantErr: true,
		},
		{
			name:    "TestNewToken-002",
			args:    args{
				userID: 1,
				req:    &request.NewToken{
					Name: "ALL",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := NewToken(tt.args.userID, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("NewToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
