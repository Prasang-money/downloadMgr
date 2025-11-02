package coordinator

import (
	"context"

	"github.com/Prasang-money/downloadMgr/internal/config"
	"github.com/Prasang-money/downloadMgr/internal/storage"
)

type Coordinator struct {
}

func NewCoordinator(store *storage.BoltStorage, nodeId string, cfg *config.Config) *Coordinator {

	return &Coordinator{}
}

func (corrd *Coordinator) StartHeartBeat(ctx context.Context) {

}

func (coord *Coordinator) StartChunkReAssignment(ctx context.Context) {

}
