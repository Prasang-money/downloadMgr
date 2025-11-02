package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Prasang-money/downloadMgr/internal/config"
	"github.com/Prasang-money/downloadMgr/internal/coordinator"
	"github.com/Prasang-money/downloadMgr/internal/downloader"
	"github.com/Prasang-money/downloadMgr/internal/storage"

	"github.com/Prasang-money/downloadMgr/internal/api"
	// "honnef.co/go/tools/config"
)

func main() {

	var (
		configPath = flag.String("config", "config.yaml", "Path to configuration file")
		nodeID     = flag.String("node-id", "", "Node identifier (auto-generated if empty)")
		apiPort    = flag.Int("port", 8080, "API server port")
		dbPath     = flag.String("db", "./data/ddm.db", "Database path")
	)
	flag.Parse()

	// Load configuration
	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Generate node ID if not provided
	if *nodeID == "" {
		*nodeID = fmt.Sprintf("node-%d", time.Now().Unix())
	}

	// Initialize storage layer
	store, err := storage.NewBoltStore(*dbPath)
	if err != nil {
		log.Fatalf("Failed to initialize storage: %v", err)
	}
	defer store.Close()

	// Initialize coordinator
	coord := coordinator.NewCoordinator(store, *nodeID, cfg)

	// Initialize downloader
	dl := downloader.NewDownloader(coord, store, cfg)

	// Start coordinator heartbeat and chunk reassignment
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go coord.StartHeartBeat(ctx)
	go coord.StartChunkReAssignment(ctx)

	// Initialize and start API server
	apiServer := api.NewServer(dl, coord, store, *apiPort)
	go func() {
		if err := apiServer.Start(); err != nil {
			log.Fatalf("API server failed: %v", err)
		}
	}()

	log.Printf("Distributed Download Manager started on node %s", *nodeID)
	log.Printf("API server listening on port %d", *apiPort)

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("Shutting down gracefully...")
	cancel()
	apiServer.Shutdown(context.Background())

}
