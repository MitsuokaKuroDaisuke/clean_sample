package interactor

import (
	"src/domain/entity"
	"src/infra/database"
	"src/usecase/interactor"
	"testing"
)

func TestUser_Create(t *testing.T) {
	type fields struct {
		UserRepository entity.UserRepository
	}
	type args struct {
		u entity.User
	}

	f := fields{
		UserRepository: &database.UserRepository{
			SQLHandler: database.NewTestSQLHandler(),
		},
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   entity.User
	}{
		{
			name:   "通常登録テスト",
			fields: f,
			args:   args{u: entity.User{Name: "UsecaseCreateUser", Email: "u-test@example.com"}},
			want:   entity.User{Name: "UsecaseCreateUser", Email: "u-test@example.com"},
		},
		{
			name:   "ポニョ吉テスト",
			fields: f,
			args:   args{u: entity.User{Name: "", Email: "u-test@example.com"}},
			want:   entity.User{Name: "ポニョ吉", Email: "u-test2@example.com"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interactor := &interactor.User{
				UserRepository: tt.fields.UserRepository,
			}
			got := interactor.Create(tt.args.u)
			if got.ID == 0 {
				t.Errorf("ユーザが作成されていません")
				return
			}
			if got.Name != tt.want.Name {
				t.Errorf("期待されるname=%v, 実際の名前=%v", tt.want.Name, got.Name)
				return
			}
		})
	}
}
