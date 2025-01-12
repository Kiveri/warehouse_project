package employees

import (
	"warehouse_project/internal/config"
	"warehouse_project/internal/domain/model"
)

type Repo struct {
	cluster      *config.Cluster
	employeesMap map[int64]*model.Employee
}

func NewRepo(cluster *config.Cluster) *Repo {
	return &Repo{
		cluster: cluster,
	}
}
