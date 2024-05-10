package boot

import (
	"github.com/thinhlh/go-market/internal/app/order/server"
	bootmanager "github.com/thinhlh/go-market/internal/core/boot_manager"
)

func Init() {
	bootmanager.Bootstrap(server.New())
}
