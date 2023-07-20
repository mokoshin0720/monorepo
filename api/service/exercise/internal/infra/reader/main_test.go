package reader_test

import (
	"os"
	"testing"

	"github.com/mokoshin0720/monorepo/app"
)

func TestMain(m *testing.M) {
	_, err := app.MySQL()
	if err != nil {
		os.Exit(1)
	}
	m.Run()
}
