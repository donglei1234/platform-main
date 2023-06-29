package redis

import (
	"go.uber.org/zap"

	"github.com/gstones/platform/services/common/nosql/memory"
	"github.com/gstones/platform/services/common/nosql/memory/redis/internal"
)

type MemoryStoreProvider = internal.MemoryStoreProvider

func NewMemoryStoreProvider(addr string, password string, l *zap.Logger) (memory.MemoryStoreProvider, error) {
	return internal.NewMemoryStoreProvider(addr, password, l)
}
