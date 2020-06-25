package tinder

import (
	p2p_discovery "github.com/libp2p/go-libp2p-discovery"
	"go.uber.org/zap"
)

// Tinder service is a simple driver backed by a cache,
type Service interface {
	Driver
}

func NewService(logger *zap.Logger, driver Driver, stratFactory p2p_discovery.BackoffFactory, opts ...p2p_discovery.BackoffDiscoveryOption) (Service, error) {
	disc, err := p2p_discovery.NewBackoffDiscovery(driver, stratFactory, opts...)
	if err != nil {
		return nil, err
	}

	return ComposeDriver("tinder", disc, disc, driver), nil
}
