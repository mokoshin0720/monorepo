package exercise

import (
	"context"

	"github.com/graph-gophers/graphql-go"
	"github.com/mokoshin0720/monorepo/app/gqlerror"
	"github.com/mokoshin0720/monorepo/service/exercise/internal/domain"
	"github.com/mokoshin0720/monorepo/service/exercise/internal/registry"
	"github.com/mokoshin0720/monorepo/service/exercise/internal/resolver"
	"github.com/mokoshin0720/monorepo/service/exercise/internal/uc"
)

// Controller Organizationのコントローラ
type Controller struct {
	registry        registry.Registry
	gqlErrorHandler gqlerror.Handler
}

// New Controllerのコンストラクタ
func New() (*Controller, error) {
	rgst, err := registry.New()
	if err != nil {
		return nil, err
	}

	h := gqlerror.Handler{
		Presenters: presenters,
	}

	return &Controller{
		registry:        rgst,
		gqlErrorHandler: h,
	}, nil
}

// Organization IDに該当するOrganizationを取得する
func (c Controller) Organization(ctx context.Context, args struct {
	ID graphql.ID
}) (resolver.Organization, error) {

	ipt := uc.GetOrganizationInput{
		ID: domain.ID(args.ID),
	}
	get := uc.NewGetOrganization(c.registry)

	opt, err := get.Do(ctx, ipt)
	if err != nil {
		return resolver.Organization{}, c.gqlErrorHandler.New(ctx, err)
	}

	r := resolver.Organization{Model: opt.Organization}
	return r, nil
}
