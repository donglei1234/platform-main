package main

import (
	"github.com/gstones/platform/services/chat/internal/app/service/external"
	"github.com/gstones/platform/services/chat/internal/app/service/public"
	"github.com/gstones/platform/services/chat/pkg/fx"
	"github.com/gstones/platform/services/common/fx/fxsvcapp"
)

func main() {
	fxsvcapp.StandardMain(
		external.ServiceModules,
		public.ServiceModule,
		fx.StoreProviderModule,
	)
}
