package application

import (
	"command-service/command/application/impl"
	"command-service/command/infra/sqlboiler"

	"go.uber.org/fx"
)

var SrvDepend = fx.Options(
	sqlboiler.RepDepend,
	fx.Provide(
		impl.NewCategoryServiceImpl,
		impl.NewProductServiceImpl,
	),
)
