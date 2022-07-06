package server

import (
	"context"
	"errors"

	"github.com/gorilla/mux"
)

type Server interface {
	Config() *Config
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("Port is required)")
	}

	if config.DatabaseUrl == "" {
		return nil, errors.New("DatabaseUrl is required")
	}

	return &Broker{
		config: config,
		router: mux.NewRouter(),
	}, nil
}
