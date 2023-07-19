package reader

import (
	"context"

	"github.com/mokoshin0720/monorepo/app"
	"github.com/mokoshin0720/monorepo/service/exercise/internal/domain"
	"github.com/mokoshin0720/monorepo/service/exercise/internal/domain/organization"
	"github.com/mokoshin0720/monorepo/service/exercise/internal/infra/adapter"
	"github.com/mokoshin0720/monorepo/service/exercise/internal/infra/entity"
	"github.com/mokoshin0720/monorepo/service/exercise/internal/infra/logger"
	"gorm.io/gorm"
)

var _ organization.Query = (*Organization)(nil)

// Organization query.Organizationの実装
type Organization struct{}

// Get 指定したIDのOrganizationを取得する。idsが空の場合は全件取得する。
// 意図せず全件取得される可能性があるため、ユースケース層でハンドリングすること。
func (o Organization) Get(ctx context.Context, tx *app.DB, ids []domain.ID) ([]organization.Organization, error) {
	orgs := []entity.Organization{}
	if err := o.preload(ctx, tx.Get()).Find(&orgs, ids).Error; err != nil {
		return []organization.Organization{}, err
	}

	return adapter.OrganizationListFromEntityList(orgs), nil
}

func (o Organization) preload(ctx context.Context, tx *gorm.DB) *gorm.DB {
	return tx.
		Session(&gorm.Session{Logger: logger.NewZerologToGormLogger(ctx)}).
		Preload("OrganizationDetail")
}
