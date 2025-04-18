package organization

import (
	"fmt"

	"github.com/mokoshin0720/monorepo/service/exercise/internal/domain"
)

const (
	maxOrganizationNameChars int = 50
)

// Name 組織名
type Name string

// Validate 組織名のバリデーション
func (n Name) Validate() error {
	if err := domain.String(n).ValidateLength(1, maxOrganizationNameChars); err != nil {
		return fmt.Errorf("alias: %w", err)
	}

	return nil
}
