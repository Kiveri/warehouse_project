package orders

import (
	"warehouse_project/internal/config"
	"warehouse_project/internal/domain/model"
)

type Repo struct {
	cluster   *config.Cluster
	ordersMap map[int64]*model.Order
}

func NewRepo(cluster *config.Cluster) *Repo {
	return &Repo{
		cluster: cluster,
	}
}
