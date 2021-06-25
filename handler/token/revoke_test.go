package token

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ismdeep/notification/api/request"
	"testing"
)

func TestRevoke(t *testing.T) {
	added, err := NewToken(1, &request.NewToken{Name: gofakeit.AppName()})
	if err != nil {
		panic(err)
	}

	type args struct {
		userID uint
		token  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestRevoke-001",
			args: args{
				userID: 1000,
				token:  "",
			},
			wantErr: true,
		},
		{
			name: "TestRevoke-002",
			args: args{
				userID: 1,
				token:  "",
			},
			wantErr: true,
		},
		{
			name: "TestRevoke-003",
			args: args{
				userID: 1,
				token:  added.Token,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Revoke(tt.args.userID, tt.args.token); (err != nil) != tt.wantErr {
				t.Errorf("Revoke() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
