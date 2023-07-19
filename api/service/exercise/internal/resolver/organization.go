package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/mokoshin0720/monorepo/service/exercise/internal/domain/organization"
)

// Organization Organizationのリゾルバ
type Organization struct {
	Model organization.Organization
}

// ID OrganizationのID
func (c Organization) ID() graphql.ID {
	return graphql.ID(c.Model.ID)
}

// Name Organizationの名前
func (c Organization) Name() string {
	return string(c.Model.Name)
}
