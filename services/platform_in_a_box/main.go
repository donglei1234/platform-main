package main

import (
	auth "github.com/gstones/platform/services/auth/pkg/module"
	"github.com/gstones/platform/services/common/fx/fxsvcapp"
	"github.com/gstones/platform/services/common/mq/mock"
	"github.com/gstones/platform/services/condition/pkg/fx"
	condition "github.com/gstones/platform/services/condition/pkg/module"
	"os"
	"path"
)

func main() {
	home := os.Getenv("HOME")
	setEnvVariable("DOCUMENT_STORE_URL", "badger://badger/"+path.Join(home, "room"))
	setEnvVariable("MEMORY_STORE_URL", "redis://127.0.0.1:6379")
	fxsvcapp.StandardMain(
		fxsvcapp.MessageQueueModule,
		mock.MQModule,
		auth.PublicModule,
		fxsvcapp.AuthStoreModule,
		fxsvcapp.ConditionStoreModule,
		condition.PublicModule,
		fx.ConditionSettingsModule,
	)
}

func setEnvVariable(key string, value string) {
	if os.Getenv(key) == "" {
		_ = os.Setenv(key, value)
	}
}
