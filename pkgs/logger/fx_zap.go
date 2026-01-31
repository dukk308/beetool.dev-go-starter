package logger

import (
	"github.com/dukk308/beetool.dev-go-starter/pkgs/logger/config"
	"go.uber.org/fx"
)

var ZapModuleFx = fx.Module("zapfx",
	fx.Provide(
		config.ProvideLogConfig,
		NewZapLogger,
		fx.Annotate(
			func(zapLogger ZapLogger) Logger {
				return zapLogger
			},
			fx.As(new(Logger))),
	),
)
