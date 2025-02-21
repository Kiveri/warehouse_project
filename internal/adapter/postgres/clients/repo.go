package clients

import (
	"errors"

	"warehouse_project/internal/config"
)

var NotFound = errors.New("client not found")

type Repo struct {
	cluster *config.Cluster
}

func NewRepo(cluster *config.Cluster) *Repo {
	return &Repo{
		cluster: cluster,
	}
}
