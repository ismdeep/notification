package auth

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ismdeep/notification/api/request"
	"github.com/ismdeep/notification/api/response"
	"github.com/ismdeep/notification/common"
	"reflect"
	"testing"
)

func TestRegister(t *testing.T) {
	type args struct {
		register *request.Register
	}
	tests := []struct {
		name    string
		args    args
		want    *response.Register
		wantErr error
	}{
		{
			name: "TestRegister-001",
			args: args{
				register: nil,
			},
			want:    nil,
			wantErr: common.ErrBadRequest,
		},
		{
			name: "TestRegister-002",
			args: args{
				register: &request.Register{
					Username: "",
					Nickname: "",
					Password: "",
				},
			},
			want:    nil,
			wantErr: common.ErrBadRequest,
		},
		{
			name: "TestRegister-003",
			args: args{
				register: &request.Register{
					Username: gofakeit.Username(),
					Nickname: "",
					Password: "",
				},
			},
			want:    nil,
			wantErr: common.ErrBadRequest,
		},
		{
			name: "TestRegister-004",
			args: args{
				register: &request.Register{
					Username: gofakeit.Username(),
					Nickname: gofakeit.Name(),
					Password: "123456",
				},
			},
			want:    &response.Register{UserID: 0},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Register(tt.args.register)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil && tt.want != nil {
				tt.want.UserID = got.UserID
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Register() got = %v, want %v", got, tt.want)
			}
		})
	}
}
