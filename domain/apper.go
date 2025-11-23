package domain

import (
	"context"
	"pdflive/config"

	"go.uber.org/zap"
)

type Apper interface {
	Options() *config.Configuration
	SetOptions(key string, value interface{}) error
	SaveOptions() error
	Logger() *zap.SugaredLogger
	ConfigPath() string
	DefaultDbPath() string
	LogPath() string
	Pwd() string
	BaseUrl() string
	Ctx() context.Context
}
