package content

import "github.com/mokoshin0720/monorepo/service/exercise/internal/domain"

type (
	// AnswerFormat 回答フォーマット
	AnswerFormat struct {
		// Vas VAS型
		Vas *VAS
		// Measure 尺度型
		Measure *Measure
	}

	// VAS VAS型
	VAS struct {
		// Min 最小値
		Min int
		// Max 最大値
		Max int
	}

	// Measure 尺度型
	Measure struct {
		// ItemList 尺度の項目のリスト
		ItemList []MeasureItem
	}

	// MeasureItem 尺度の項目
	MeasureItem struct {
		// Level 尺度のレベル
		Level int
		// Text 尺度のテキスト
		Text domain.String
	}
)
