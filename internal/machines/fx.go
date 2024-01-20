package machines

import "go.uber.org/fx"

var Module = fx.Module("machines",
	fx.Provide(New),
)
