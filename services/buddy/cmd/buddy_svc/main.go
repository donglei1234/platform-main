package main

import (
	"github.com/gstones/platform/services/buddy/internal/app/service/external"
	"github.com/gstones/platform/services/buddy/internal/app/service/private"
	"github.com/gstones/platform/services/buddy/internal/app/service/public"
	"github.com/gstones/platform/services/common/fx/fxsvcapp"
)

func main() {
	fxsvcapp.StandardMain(
		external.ServiceModules,
		public.ServiceModule,
		private.ServiceModule,
	)
}
