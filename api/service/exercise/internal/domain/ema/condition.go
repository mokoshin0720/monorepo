package ema

// Condition EMAの回答における状態を表す
type Condition int

// Validate Conditionのバリデーションを行う
func (c Condition) Validate() error {
	// TODO: 最小値と最大値のチェック
	return nil
}
