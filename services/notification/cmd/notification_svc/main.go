package main

import (
	"math/rand"
	"os"
	"path"
	"time"

	afx "github.com/gstones/platform/services/auth/pkg/fx"
	auth "github.com/gstones/platform/services/auth/pkg/module"
	"github.com/gstones/platform/services/common/fx/fxsvcapp"
	"github.com/gstones/platform/services/notification/internal/app/service/public"
	ffx "github.com/gstones/platform/services/notification/pkg/fx"
)

func main() {
	rand.Seed(time.Now().Unix())
	home := os.Getenv("HOME")
	setEnvVariable("DOCUMENT_STORE_URL", "badger://badger/"+path.Join(home, "notification"))

	fxsvcapp.StandardMain(
		ffx.NotificationSettingsModule,
		public.ServiceModule,
		auth.PublicModule,
		afx.AuthSettingsModule,
		fxsvcapp.AuthStoreModule,
		fxsvcapp.AuthClientModule,
	)
}

func setEnvVariable(key string, value string) {
	if os.Getenv(key) == "" {
		_ = os.Setenv(key, value)
	}
}
