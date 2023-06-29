package main

import (
	"math/rand"
	"os"
	"path"
	"time"

	auth "github.com/gstones/platform/services/auth/pkg/module"
	"github.com/gstones/platform/services/common/fx/fxsvcapp"
	"github.com/gstones/platform/services/guild/internal/app/service/external"
	"github.com/gstones/platform/services/guild/internal/app/service/public"
	ffx "github.com/gstones/platform/services/guild/pkg/fx"
)

func main() {
	rand.Seed(time.Now().Unix())
	home := os.Getenv("HOME")
	setEnvVariable("DOCUMENT_STORE_URL", "badger://badger/"+path.Join(home, "notification"))

	fxsvcapp.StandardMain(
		ffx.GuildSettingsModule,
		external.ServicesModules,
		public.ServiceModule,
		fxsvcapp.AuthStoreModule,
		auth.PublicModule,
	)
}

func setEnvVariable(key string, value string) {
	if os.Getenv(key) == "" {
		_ = os.Setenv(key, value)
	}
}
