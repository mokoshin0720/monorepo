package user

import (
	"time"

	"github.com/mokoshin0720/monorepo/service/exercise/internal/domain"
	"github.com/mokoshin0720/monorepo/service/exercise/internal/domain/ema"
)

type (
	// User ユーザーを表現するドメインモデルの集約
	User struct {
		// ID ユーザーのID
		ID UserID
		// Name 名前
		Name domain.String
		// Type ユーザーのタイプ
		Type Type
		// IsAnsweredEMA EMAに回答済みかどうか
		IsAnsweredEMA bool
		// RegistrationDate 登録日時
		RegistrationDate time.Time
		// EMAList 回答したEMAのリスト
		EMAList []ema.EMA
	}

	UserID domain.ID
)

// Validate ユーザーのバリデーションを行う
func (u User) Validate() error {
	// TODO: 実際のmin/maxに変更する
	if err := u.Name.ValidateLength(0, 0); err != nil {
		return err
	}
	return nil
}
