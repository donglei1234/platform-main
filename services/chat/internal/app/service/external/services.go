package external

import (
	cfx "github.com/gstones/platform/services/chat/pkg/fx"
	"github.com/gstones/platform/services/common/fx/fxsvcapp"
	"github.com/gstones/platform/services/common/mq/nats"
	"go.uber.org/fx"
)

var ServiceModules = fx.Options(
	fxsvcapp.MessageQueueModule,
	nats.MQModule,
	cfx.ChatSettingsModule,
)
