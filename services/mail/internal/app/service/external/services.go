package external

import (
	"go.uber.org/fx"

	"github.com/gstones/platform/services/common/fx/fxsvcapp"
)

var ServicesModules = fx.Options(
	fxsvcapp.RoomStoreModule,
)
