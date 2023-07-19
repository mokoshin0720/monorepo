package domain_test

import (
	"testing"

	"github.com/mokoshin0720/monorepo/service/exercise/internal/domain"
)

func Test_NewID(t *testing.T) {
	t.Parallel()

	id := domain.NewID()
	t.Log(id)
}
