package operator

import (
	"fmt"

	"github.com/dgraph-io/badger/v2"
	"github.com/rs/zerolog"
)

type BadgerToGlobalLogger struct {
	badger.Logger

	logger *zerolog.Logger
}

func NewBadgerLogConverter(logger zerolog.Logger) *BadgerToGlobalLogger {
	return &BadgerToGlobalLogger{
		logger: &logger,
	}
}

func (b *BadgerToGlobalLogger) Infof(str string, i ...interface{}) {
	b.logger.Info().Msg(fmt.Sprintf(str, i...))
}

func (b *BadgerToGlobalLogger) Warningf(str string, i ...interface{}) {
	b.logger.Warn().Msg(fmt.Sprintf(str, i...))
}

func (b *BadgerToGlobalLogger) Debugf(str string, i ...interface{}) {
	b.logger.Debug().Msg(fmt.Sprintf(str, i...))
}

func (b *BadgerToGlobalLogger) Errorf(str string, i ...interface{}) {
	b.logger.Error().Msg(fmt.Sprintf(str, i...))
}
