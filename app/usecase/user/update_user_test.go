// これは未実装です

package user

import (
    "context"
    "errors"
    "reflect"
    "testing"

    gomock "go.uber.org/mock/gomock"
    userDomain "github.com/onion0904/app/domain/user"
)

func TestUpdateUseCaseRun(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()
    mockRepo := userDomain.NewMockUserRepository(ctrl)

    existingUser, _ := userDomain.Reconstruct(
        "123", "Doe", "John", "john@example.com", 
        "hashedpassword", "icon.png", 
		nil,nil,
    )
    
    updatedUser, _ := userDomain.Reconstruct(
        "123", "Updated", "User", "updated@example.com", 
        "hashedpassword", "Updated Icon", 
		nil,nil,
    )

    tests := []struct {
        name    string
        setup   func()
        args    struct {
            ctx context.Context
            id  string
            dto UpdateUseCaseDto
        }
        want    *userDomain.User
        wantErr bool
    }{
        {
            name: "成功: ユーザー更新",
            setup: func() {
                // FindUser が既存のユーザーを返すように設定（エラーなし）
                mockRepo.EXPECT().
                    FindUser(gomock.Any(), "123").
                    Return(existingUser, nil).
                    Times(1)

                // Update が正常に動作することを確認
                mockRepo.EXPECT().
                    Update(gomock.Any(), gomock.Any()).
                    Return(nil).
                    Times(1)

                // 再度 FindUser を呼び出した時に更新後のユーザーを返す
                mockRepo.EXPECT().
                    FindUser(gomock.Any(), "123").
                    Return(updatedUser, nil).
                    Times(1)
            },
            args: struct {
                ctx context.Context
                id  string
                dto UpdateUseCaseDto
            }{
                ctx: context.Background(),
                id:  "123",
                dto: UpdateUseCaseDto{
                    LastName:  "Updated",
                    FirstName: "User",
                    Email:     "updated@example.com",
                    Icon:      "Updated Icon",
                },
            },
            want:    updatedUser,
            wantErr: false,
        },
        {
            name: "失敗: ユーザーが見つからない",
            setup: func() {
                // FindUser がエラーを返すように設定
                mockRepo.EXPECT().
                    FindUser(gomock.Any(), "456").
                    Return(nil, errors.New("user not found")).
                    Times(1)
            },
            args: struct {
                ctx context.Context
                id  string
                dto UpdateUseCaseDto
            }{
                ctx: context.Background(),
                id:  "456",
                dto: UpdateUseCaseDto{
                    LastName:  "Fail",
                    FirstName: "Test",
                    Email:     "fail@example.com",
                },
            },
            want:    nil,
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            tt.setup()

            uc := NewUpdateUserUseCase(mockRepo)

            got, err := uc.Run(tt.args.ctx, tt.args.id, tt.args.dto)

            if (err != nil) != tt.wantErr {
                t.Errorf("UpdateUseCase.Run() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("UpdateUseCase.Run() = %v, want %v", got, tt.want)
            }
        })
    }
}