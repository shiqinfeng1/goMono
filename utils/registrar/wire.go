package registrar

import (
	"github.com/google/wire"
)

// ProviderSet is server providers.  使用哪个就注册哪个
var ProviderSet = wire.NewSet(MustNacosRegistrar)
