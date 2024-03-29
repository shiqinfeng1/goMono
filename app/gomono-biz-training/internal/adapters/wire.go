package adapters

import (
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewTrainingRepo, NewTrainerGrpc, NewUserGrpc)
