package mail

import (
	"github.com/ismdeep/notification/api/request"
	"testing"
)

func TestPushMail(t *testing.T) {
	type args struct {
		userID uint
		req    *request.PushMailRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			// req is nil
			name: "",
			args: args{
				userID: 0,
				req:    nil,
			},
			wantErr: true,
		},
		{
			// bad request with check()
			name: "",
			args: args{
				userID: 1,
				req: &request.PushMailRequest{
					SenderName: "",
					Subject:    "",
					Type:       "",
					Content:    "",
					ToMailList: nil,
				},
			},
			wantErr: true,
		},
		{
			// 用户不存在
			name: "",
			args: args{
				userID: 10000,
				req: &request.PushMailRequest{
					SenderName: "JustOJ",
					Subject:    "Hello",
					Type:       "text/plain",
					Content:    "Hi, there.",
					ToMailList: []string{"l.jiang.1024@gmail.com"},
				},
			},
			wantErr: true,
		},
		{
			// type错误
			name: "",
			args: args{
				userID: 1,
				req: &request.PushMailRequest{
					SenderName: "JustOJ",
					Subject:    "Hello",
					Type:       "invalid type",
					Content:    "Hi, there.",
					ToMailList: []string{"l.jiang.1024@gmail.com"},
				},
			},
			wantErr: true,
		},
		{
			// 收件人列表为空
			name: "",
			args: args{
				userID: 1,
				req: &request.PushMailRequest{
					SenderName: "JustOJ",
					Subject:    "Hello",
					Type:       "text/plain",
					Content:    "Hi, there.",
					ToMailList: nil,
				},
			},
			wantErr: true,
		},
		{
			// 正常请求
			name: "",
			args: args{
				userID: 1,
				req: &request.PushMailRequest{
					SenderName: "JustOJ",
					Subject:    "Hello",
					Type:       "text/plain",
					Content:    "Hi, there.",
					ToMailList: []string{"l.jiang.1024@gmail.com"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PushMail(tt.args.userID, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("PushMail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
