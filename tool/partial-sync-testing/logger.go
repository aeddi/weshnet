package main

import (
	"berty.tech/go-orbit-db/stores"

	"go.uber.org/zap"
)

func newLogger(level string) *zap.Logger {
	var err error

	cfg := zap.NewDevelopmentConfig()
	cfg.Level, err = zap.ParseAtomicLevel(level)
	if err != nil {
		panic(err)
	}

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	return logger
}

func logEvent(logger *zap.Logger, evt interface{}) {
	switch e := evt.(type) {
	case stores.EventReady:
		logger.Info("EventReady",
			zap.String("Address", e.Address.String()),
			zap.Int("Heads length", len(e.Heads)),
		)
	case stores.EventNewPeer:
		logger.Info("EventNewPeer",
			zap.String("Peer ID", e.Peer.String()),
		)
	case stores.EventLoad:
		logger.Info("EventLoad",
			zap.String("Address", e.Address.String()),
			zap.Int("Heads length", len(e.Heads)),
		)
	case stores.EventLoadProgress:
		logger.Info("EventLoadProgress",
			zap.String("Address", e.Address.String()),
			zap.String("Entry", e.Entry.GetHash().String()),
		)
	case stores.EventReplicate:
		logger.Info("EventReplicate",
			zap.String("Address", e.Address.String()),
		)
	case stores.EventReplicateProgress:
		logger.Info("EventReplicateProgress",
			zap.String("Address", e.Address.String()),
			zap.Int("Max", e.Max),
			zap.Int("Progress", e.Progress),
			zap.String("Entry", e.Entry.GetHash().String()),
		)
	case stores.EventReplicated:
		logger.Info("EventReplicated",
			zap.String("Address", e.Address.String()),
			zap.Int("Log length", e.LogLength),
		)
	case stores.EventWrite:
		logger.Info("EventWrite",
			zap.String("Address", e.Address.String()),
			zap.String("Entry", e.Entry.GetHash().String()),
			zap.Int("Heads length", len(e.Heads)),
		)
	}
}
