package session

import (
	"github.com/fulgurant/datastore"
	"go.uber.org/zap"
)

type Options struct {
	config    *Config
	logger    *zap.Logger
	datastore datastore.Handler
}

func DefaultOptions(config *Config) *Options {
	return &Options{
		config: config,
	}
}

// WithLogger sets the logger for the server
func (o *Options) WithLogger(value *zap.Logger) *Options {
	o.logger = value.With(zap.String("module", "session-handler"))
	return o
}

// WithDataStore sets the datastore for the server
func (o *Options) WithDataStore(value datastore.Handler) *Options {
	o.datastore = value
	return o
}
