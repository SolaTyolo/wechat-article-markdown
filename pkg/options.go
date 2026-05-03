package kbsink

import (
	"net/http"

	"github.com/kbsink-org/kbsink/pkg/core"
)

type converterConfig struct {
	driver core.Driver
	parser core.Parser
	store  core.Storage
	client *http.Client
}

// Option configures a Converter.
type Option func(*converterConfig)

func WithDriver(d core.Driver) Option {
	return func(c *converterConfig) { c.driver = d }
}

func WithParser(p core.Parser) Option {
	return func(c *converterConfig) { c.parser = p }
}

func WithStorage(s core.Storage) Option {
	return func(c *converterConfig) { c.store = s }
}

func WithHTTPClient(client *http.Client) Option {
	return func(c *converterConfig) { c.client = client }
}
