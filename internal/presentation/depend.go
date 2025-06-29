package presentation

import (
	"command-service/internal/application"
	"command-service/internal/presentation/adapter"
	"command-service/internal/presentation/prepare"
	"command-service/internal/presentation/server"

	"go.uber.org/fx"
)

var CommandDepend = fx.Options(
	application.SrvDepend,
	fx.Provide(
		adapter.NewCategoryAdapterImpl,
		adapter.NewProductAdapterImpl,
		server.NewCategoryServer,
		server.NewProductServer,
		prepare.NewCommandServer,
	),
	fx.Invoke(prepare.CommandServiceLifecycle),
)
