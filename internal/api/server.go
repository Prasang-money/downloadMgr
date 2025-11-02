package api

import (
	"context"

	"github.com/Prasang-money/downloadMgr/internal/coordinator"
	"github.com/Prasang-money/downloadMgr/internal/downloader"
	"github.com/Prasang-money/downloadMgr/internal/storage"
)

type Server struct {
}

func NewServer(dl *downloader.Downloader, coord *coordinator.Coordinator, store *storage.BoltStorage, port int) *Server {

	return &Server{}
}

func (server *Server) Start() error {

	return nil
}

func (server *Server) Shutdown(ctx context.Context) {

}
