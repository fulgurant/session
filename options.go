package session

import (
	"errors"

	"github.com/fulgurant/datastore"
	"go.uber.org/zap"
)

type Options struct {
	config    *Config
	logger    *zap.Logger
	datastore datastore.GetSetter
}

func DefaultOptions(config *Config) (*Options, error) {
	if config == nil {
		return nil, errors.New("nil config not allowed")
	}

	return &Options{
		config: config,
	}, nil
}

// WithLogger sets the logger for the server
func (o *Options) WithLogger(value *zap.Logger) *Options {
	o.logger = value.With(zap.String("module", "session-handler"))
	return o
}

// WithDataStore sets the datastore for the server
func (o *Options) WithDataStore(value *datastore.GetSetter) *Options {
	o.datastore = value
	return o
}
