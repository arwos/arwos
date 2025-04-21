package settings

import (
	"context"

	"go.osspkg.com/goppy/v2/orm"
)

type Service struct {
	repo *repoModels
}

func NewService(db orm.ORM) *Service {
	return &Service{repo: newRepoModels(db)}
}

func (s *Service) GetEnvByPluginId(ctx context.Context, id int64) ([]EnvModel, error) {
	return s.repo.ReadEnvModelByPluginId(ctx, id)
}
