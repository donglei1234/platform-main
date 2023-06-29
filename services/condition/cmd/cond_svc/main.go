package main

import (
	auth "github.com/gstones/platform/services/auth/pkg/module"
	"github.com/gstones/platform/services/common/fx/fxsvcapp"
	"github.com/gstones/platform/services/common/mq/mock"
	"github.com/gstones/platform/services/condition/internal/app/service/external"
	"github.com/gstones/platform/services/condition/internal/app/service/public"
	"github.com/gstones/platform/services/condition/pkg/fx"
	"os"
	"path"
)

func main() {
	home := os.Getenv("HOME")
	setEnvVariable("DOCUMENT_STORE_URL", "badger://badger/"+path.Join(home, "room"))
	setEnvVariable("MEMORY_STORE_URL", "redis://127.0.0.1:6379/0")

	fxsvcapp.StandardMain(
		fx.ConditionSettingsModule,
		public.ServiceModule,
		external.ServicesModules,
		auth.PublicModule,
		fxsvcapp.AuthStoreModule,
		fxsvcapp.MessageQueueModule,
		mock.MQModule,
	)
}

func setEnvVariable(key string, value string) {
	if os.Getenv(key) == "" {
		_ = os.Setenv(key, value)
	}
}
