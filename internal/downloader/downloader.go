package downloader

import (
	"github.com/Prasang-money/downloadMgr/internal/config"
	"github.com/Prasang-money/downloadMgr/internal/coordinator"
	"github.com/Prasang-money/downloadMgr/internal/storage"
)

type Downloader struct {
}

func NewDownloader(coord *coordinator.Coordinator, store *storage.BoltStorage, cfg *config.Config) *Downloader {

	return &Downloader{}
}
