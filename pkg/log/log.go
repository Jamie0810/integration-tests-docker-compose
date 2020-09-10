package log

import (
	"fmt"
	"os"
	"strings"
	"time"

	cnf "gitlab.silkrode.com.tw/team_golang/kbc/integration-tests/configs"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

// NewLogger returns the logger
func InitLogger(cfg cnf.Config) (logger *Logger, err error) {
	level := zerolog.InfoLevel
	level, err = zerolog.ParseLevel(strings.ToLower(cfg.LogLevel))
	if err != nil {
		return nil, err
	}
	zerolog.SetGlobalLevel(level)
	zerolog.TimeFieldFormat = time.RFC3339Nano
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	switch strings.ToLower(cfg.LogFormat) {
	case "json":
		logger = &Logger{zerolog.New(os.Stdout).With().Timestamp().Logger()}
	default:
		logger = &Logger{log.Output(
			zerolog.ConsoleWriter{
				Out:        os.Stdout,
				TimeFormat: zerolog.TimeFieldFormat,
			}),
		}
	}

	return
}

//Logger wrapper for Zerolog
type Logger struct{ zerolog.Logger }

//TraceMsg output tracelevel msg
func (l *Logger) TraceMsg(a ...interface{}) {
	log := l.With().CallerWithSkipFrameCount(3).Logger()
	log.Trace().Msg(fmt.Sprint(a...))
}

//DebugMsg output debuglevel msg
func (l *Logger) DebugMsg(a ...interface{}) {
	log := l.With().CallerWithSkipFrameCount(3).Logger()
	log.Debug().Msg(fmt.Sprint(a...))
}

//InfoMsg output infolevel msg
func (l *Logger) InfoMsg(a ...interface{}) {
	log := l.With().CallerWithSkipFrameCount(3).Logger()
	log.Info().Msg(fmt.Sprint(a...))
}

//WarnMsg output warnlevel msg
func (l *Logger) WarnMsg(a ...interface{}) {
	log := l.With().CallerWithSkipFrameCount(3).Logger()
	log.Warn().Msg(fmt.Sprint(a...))
}

//ErrorMsg output errorlevel msg
func (l *Logger) ErrorMsg(a ...interface{}) {
	log := l.With().CallerWithSkipFrameCount(3).Logger()
	log.Error().Msg(fmt.Sprint(a...))
}

//FatalMsg output fatallevel msg
func (l *Logger) FatalMsg(a ...interface{}) {
	log := l.With().CallerWithSkipFrameCount(3).Logger()
	log.Fatal().Msg(fmt.Sprint(a...))
}

//PanicMsg output paniclevel msg
func (l *Logger) PanicMsg(a ...interface{}) {
	log := l.With().CallerWithSkipFrameCount(3).Logger()
	log.Panic().Msg(fmt.Sprint(a...))
}
