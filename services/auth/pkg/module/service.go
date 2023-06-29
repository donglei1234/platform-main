package module

import (
	"github.com/gstones/platform/services/auth/internal/app/service/public"
	"github.com/gstones/platform/services/auth/pkg/metadata"
	"go.uber.org/fx"
)

var (
	Modules = fx.Module(metadata.AppId,
		public.ServiceModule,
	)
)
