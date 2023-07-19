package content

import "github.com/mokoshin0720/monorepo/service/exercise/internal/domain"

type (
	// Chapter 章を表現するドメインモデルの集約
	Chapter struct {
		// ID 章のID
		ID ChapterID
		// Title 章のタイトル
		Title domain.String
		// EpisodeList チャプターに含まれるエピソードのリスト
		EpisodeList []Episode
	}

	// ChapterID 章のID
	ChapterID domain.ID
)
