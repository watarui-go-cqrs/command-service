package sqlboiler

import (
	"command-service/internal/infrastructure/sqlboiler/repository"

	"go.uber.org/fx"
)

var RepDepend = fx.Options(
	fx.Provide(
		repository.NewCategoryRepositorySQLBoiler,
		repository.NewProductRepositorySQLBoiler,
	),
)
