package content

import "github.com/mokoshin0720/monorepo/service/exercise/internal/domain"

type (
	// Block ブロックを表現するドメインモデルの集約
	Block struct {
		// ID ブロックのID
		ID BlockID
		// WorkBlock ワークを表現するブロック
		WorkBlock *WorkBlock
		// ImageBlock 画像を表現するブロック
		ImageBlock *ImageBlock
		// VideoBlock 動画を表現するブロック
		VideoBlock *VideoBlock
		// SoundBlock 音声を表現するブロック
		SoundBlock *SoundBlock
		// TalkBlock 会話を表現するブロック
		TalkBlock *TalkBlock
	}

	// WorkBlock ワークを表現するブロック
	WorkBlock struct {
		// AnswerFormatList 回答フォーマットのリスト
		AnswerFormatList []AnswerFormat
	}

	// ImageBlock 画像を表現するブロック
	ImageBlock struct {
		// ImageURL 画像のURL
		ImageURL domain.String
	}

	// VideoBlock 動画を表現するブロック
	VideoBlock struct {
		// VideoURL 動画のURL
		VideoURL domain.String
	}

	// SoundBlock 音声を表現するブロック
	SoundBlock struct {
		// SoundURL 音声のURL
		SoundURL domain.String
	}

	// TalkBlock 会話を表現するブロック
	TalkBlock struct {
		// Message 会話のメッセージ
		Message domain.String
	}

	// BlockID ブロックのID
	BlockID domain.ID
)
