package exercise

import (
	"errors"

	"github.com/mokoshin0720/monorepo/app/gqlerror"
	"golang.org/x/text/language"
)

var (
	// ErrUnauthorized 認証に失敗したエラー
	ErrUnauthorized = errors.New("unauthorized")
)

var presenters = map[error]gqlerror.Presenter{
	ErrUnauthorized: {
		Code: "unauthorized",
		Lang2Msg: map[language.Tag]string{
			language.Japanese: "リクエストが無効です。",
			language.English:  "Invalid request",
		},
	},
}
