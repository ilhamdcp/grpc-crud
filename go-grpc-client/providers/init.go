package providers

import "github.com/google/wire"

var SuperSet = wire.NewSet(ProvideClientConnection)
