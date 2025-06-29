package application

import (
	"command-service/internal/application/impl"
	"command-service/internal/infrastructure/sqlboiler"

	"go.uber.org/fx"
)

var SrvDepend = fx.Options(
	sqlboiler.RepDepend,
	fx.Provide(
		impl.NewCategoryServiceImpl,
		impl.NewProductServiceImpl,
	),
)
