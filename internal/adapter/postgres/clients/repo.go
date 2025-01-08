package clients

import (
	"warehouse_project/internal/config"
	"warehouse_project/internal/domain/model"
)

type Repo struct {
	cluster    *config.Cluster
	clientsMap map[int64]*model.Client
}

func NewRepo(cluster *config.Cluster) *Repo {
	return &Repo{
		cluster: cluster,
	}
}
