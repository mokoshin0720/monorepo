package ema

import (
	"time"

	"github.com/mokoshin0720/monorepo/service/exercise/internal/domain"
)

type (
	// EMA EMAを表現するドメインモデルの集約
	EMA struct {
		// ID EMAのID
		ID EMAID
		// Day 回答した日付
		Day time.Time
		// Tiredness 疲労度
		Tiredness Condition
		// TODO: 他の項目を追加する
	}

	// EMAID EMAのIDZ
	EMAID domain.ID
)
