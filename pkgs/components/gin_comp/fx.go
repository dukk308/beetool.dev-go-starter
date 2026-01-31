package gin_comp

import (
	"github.com/dukk308/beetool.dev-go-starter/pkgs/global_config"
	"go.uber.org/fx"
)

func ProvideGinConfig(global_config *global_config.GlobalConfig) *GinConfig {
	return LoadGinConfig(global_config)
}

var GinComponentFx = fx.Module("gin",
	fx.Provide(ProvideGinConfig),
	fx.Provide(NewGinComp),
)
