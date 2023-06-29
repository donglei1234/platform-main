package external

import (
	"go.uber.org/fx"

	ffx "github.com/gstones/platform/services/buddy/pkg/fx"
)

var ServiceModules = fx.Options(
	ffx.BuddySettingsModule,
)
