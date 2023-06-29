package module

import (
	"github.com/gstones/platform/services/common/fx/fxsvcapp"
	"github.com/gstones/platform/services/condition/internal/app/service/public"
	fx2 "github.com/gstones/platform/services/condition/pkg/fx"
	"github.com/gstones/platform/services/condition/pkg/metadata"
	"go.uber.org/fx"
)

var (
	PublicModule = fx.Module(
		metadata.AppId,
		public.ServiceModule,
		fx2.ConditionClientModule,
		fx2.ConditionSettingsModule,
		fxsvcapp.ConditionStoreModule,
	)
)
