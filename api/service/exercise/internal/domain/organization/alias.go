package organization

import (
	"fmt"

	"github.com/mokoshin0720/monorepo/service/exercise/internal/domain"
)

const (
	maxOrganizationAliasChars int = 50
)

// Alias 組織の略称
type Alias domain.String

// Validate 組織の略称のバリデーション
func (a Alias) Validate() error {
	if err := domain.String(a).ValidateLength(1, maxOrganizationAliasChars); err != nil {
		return fmt.Errorf("alias: %w", err)
	}

	return nil
}
