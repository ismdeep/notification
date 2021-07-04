package auth

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ismdeep/notification/api/request"
	"github.com/ismdeep/notification/api/response"
	"reflect"
	"testing"
)

func TestLogin(t *testing.T) {
	username := gofakeit.Username()
	password := gofakeit.Password(true, true, true, false, false, 10)

	added, err := Register(&request.Register{
		Username: username,
		Nickname: gofakeit.Name(),
		Password: password,
	})
	if err != nil {
		panic(err)
	}

	if added == nil {
		panic("added is nil")
	}

	type args struct {
		login *request.Login
	}
	tests := []struct {
		name    string
		args    args
		want    *response.Login
		wantErr bool
	}{
		{
			name: "TestLogin-001",
			args: args{
				login: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "TestLogin-002",
			args: args{
				login: &request.Login{
					Username: "",
					Password: "",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "TestLogin-003",
			args: args{
				login: &request.Login{
					Username: username,
					Password: password,
				},
			},
			want: &response.Login{
				UserID:      added.UserID,
				AccessToken: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Login(tt.args.login)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil && tt.want != nil {
				tt.want.AccessToken = got.AccessToken
				t.Logf("access_token = %v", got.AccessToken)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Login() got = %v, want %v", got, tt.want)
			}
		})
	}
}
