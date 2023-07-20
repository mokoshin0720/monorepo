package content

import "github.com/mokoshin0720/monorepo/service/exercise/internal/domain"

type (
	// Episode エピソードを表現するドメインモデルの集約
	Episode struct {
		// ID エピソードのID
		ID EpisodeID
		// Title エピソードのタイトル
		Title domain.String
	}

	// EpisodeID エピソードのID
	EpisodeID int
)
