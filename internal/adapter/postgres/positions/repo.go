package positions

import (
	"warehouse_project/internal/config"
	"warehouse_project/internal/domain/model"
)

type Repo struct {
	cluster      *config.Cluster
	positionsMap map[int64]*model.Position
}

func NewRepo(cluster *config.Cluster) *Repo {
	return &Repo{
		cluster: cluster,
	}
}
