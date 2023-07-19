package main

import (
	"github.com/mokoshin0720/monorepo/app"
	"github.com/mokoshin0720/monorepo/service/exercise/seed"
	"gorm.io/gorm"
)

func main() {
	db, err := app.MySQL()

	if err != nil {
		panic(err)
	}

	seeds := seed.Dev()

	err = db.Transaction(func(tx *gorm.DB) error {
		for _, s := range seeds {
			if err := tx.Create(s).Error; err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		panic(err)
	}
}
