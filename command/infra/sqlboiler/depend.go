package sqlboiler

import (
	"command-service/command/infra/sqlboiler/repository"

	"go.uber.org/fx"
)

var RepDepend = fx.Options(
	fx.Provide(
		repository.NewCategoryRepositorySQLBoiler,
		repository.NewProductRepositorySQLBoiler,
	),
)
